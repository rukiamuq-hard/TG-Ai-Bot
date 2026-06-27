package main

import (
	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
	"log"
	"main.go/ai"
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

	b.Handle("/Claude", func(c tele.Context) error {
		tags := c.Args()
		if tags == nil {
			return c.Reply(c.Send("Wrong format\nUse: /command <prompt>\n"))
		}
		s, sep := "", " "
		for _, tag := range tags {
			s += tag + sep

		}
		response, err := ai.SendMessageToClaude(s)
		if err != nil {
			return c.Reply(c.Send(err))
		}
		return c.Reply(c.Send(response))
	})

	/*
		b.Handle(tele.OnText, func(c tele.Context) error {
			sender := c.Sender()
			fmt.Println("Sender: ", sender, "\nText: ", c.Text())
			return nil
		})
	*/
	b.Start()
}
