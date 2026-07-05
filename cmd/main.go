package main

import (
	"TgAiBot/ai"
	"TgAiBot/dataBase/context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
	_ "modernc.org/sqlite"
)

func init() {
	if err := godotenv.Load("../TokensChatId.env"); err != nil {
		log.Fatal("Error load env: ", err)
	}
}

func main() {
	token := os.Getenv("TOKEN")

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 5 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	prompt, err := os.ReadFile("prompt.txt")
	if err != nil {
		log.Println("Default working, without prompt", err)
	}

	DB, err := dataBaseContext.CreateStartDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dataBaseContext.CloseDB(DB)

	b.Handle("/Gemini", func(c tele.Context) error {
		if c.Args() == nil {
			return c.Reply("Wrong format\nUse: /Gemini <prompt>\n")
		}

		s := strings.Join(c.Args(), " ")

		resp, err := dataBaseContext.ReadFromContextDB(DB, c.Sender().ID)
		if err != nil {
			log.Println("Error: ", err)
			return c.Reply("Error, try again")
		}

		response, err := ai.GeminiGetResponse(resp, "[PROMPT]: "+string(prompt)+"[QUESTION]: "+s)
		if err != nil {
			log.Println("Error: ", err)
			return c.Reply("Error, try again")
		}

		err = dataBaseContext.StoreToContextDB(DB, c.Sender().ID, "user", s)
		if err != nil {
			log.Println("Error: ", err)
			return c.Reply("Error, try again")
		}
		err = dataBaseContext.StoreToContextDB(DB, c.Sender().ID, "model", response)
		if err != nil {
			log.Println("Error: ", err)
			return c.Reply("Error, try again!")
		}
		return c.Reply(response)
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		err := dataBaseContext.StoreToChatLogDB(DB, c.Sender().FirstName+" "+c.Sender().LastName, c.Text())
		if err != nil {
			log.Fatal(err)
		}
		return nil
	})

	b.Handle("/ChatLogs", func(c tele.Context) error {

		for _, s := range c.Message().Payload {
			if unicode.IsLetter(s) || unicode.IsSymbol(s) {
				return c.Reply("Wrong format. Use numbers")
			}
		}

		val, _ := strconv.ParseInt(c.Message().Payload, 10, 64)
		if val == 0 {
			val = 200
		}

		str, err := dataBaseContext.ReadFromChatLogDB(DB, val)
		if err != nil {
			log.Println("Error: ", err)
			return c.Reply("Error, try again later.")
		}
		response, err := ai.GeminiGetResponseNoHistory("(Всё что в скобках-системный промпт, твоя цель пересказать коротко что произошло в этих сообщениях, от самых левых, тоесть новых, к старым справа): " + str)
		if err != nil {
			log.Println("Error: ", err)
			return c.Reply("Error, try again later.")
		}
		return c.Reply(response + " " + c.Message().Payload)
	})
	b.Start()
}
