package handler

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func (h *Handler) GeminiGetResp(c tele.Context) error {
	if c.Args() == nil {
		return c.Reply("Wrong format\nUse: /Gemini <prompt>\n")
	}

	s := strings.Join(c.Args(), " ")

	resp, err := database.ReadFromContextDB(DB, c.Sender().ID)
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
}

func (h *Handler) ChatLogs(c tele.Context) error {
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
}
