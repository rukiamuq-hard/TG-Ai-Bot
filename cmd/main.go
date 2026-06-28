package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
	"main.go/ai"
	"main.go/dataBase"
	_ "modernc.org/sqlite"
)

func init() {
	if err := godotenv.Load("../TokensChatId.env"); err != nil {
		log.Print("Error load .env")
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

	storage, err := database.CreateStartDB()
	if err != nil {
		log.Fatal(err)
	}
	defer database.CloseDB(storage)

	b.Handle("/Gemini", func(c tele.Context) error {
		if c.Args() == nil {
			return c.Reply("Wrong format\nUse: /Gemini <prompt>\n")
		}

		s := strings.Join(c.Args(), " ")

		resp := database.ReadFromDB(storage, c.Sender().ID)
		response, err := ai.SendMessageToGemini("[HISTORY]: " + resp + "[QUESTION]: " + s)
		if err != nil {
			return c.Reply(err)
		}
		go database.InsertToDB(storage, c.Sender().ID, s)
		return c.Reply(response)
	})
	//This for logs of the chat
	/*
		b.Handle(tele.OnText, func(c tele.Context) error {
			sender := c.Sender()
			fmt.Println("Sender: ", sender, "\nText: ", c.Text())
			return nil
		})
	*/
	b.Start()
}
