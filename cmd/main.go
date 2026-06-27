package main

import (
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
	"log"
	"main.go/ai"
	"main.go/dataBase"
	_ "modernc.org/sqlite"
	"os"
	"time"
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
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
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
		tags := c.Args()
		if tags == nil {
			return c.Reply("Wrong format\nUse: /Gemini <prompt>\n")
		}

		var s string
		for _, tag := range tags {
			s += tag + " "
		}
		resp := database.ReadFromDB(storage, c.Sender().ID)

		response, err := ai.SendMessageToGemini(resp + "Вопрос: " + s)
		if err != nil {
			return c.Reply(err)
		}
		database.InsertToDB(storage, c.Sender().ID, s)
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
