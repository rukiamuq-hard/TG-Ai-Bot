package handler

import (
	"context"
	tele "gopkg.in/telebot.v4"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func (h *Handler) GeminiGetResp(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if c.Args() == nil {
		return c.Reply("Usage: /Gemini <prompt>")
	}

	s := strings.Join(c.Args(), " ")

	resp, err := h.service.ServiceReadFromContextDB(ctx, c.Sender().ID)
	if err != nil {
		log.Println("Error: ", err)
		return c.Reply("Error, try again")
	}

	response, err := h.service.ServiceGeminiGetResponse(ctx, resp, "[QUESTION]: "+s)
	if err != nil {
		log.Println("Error: ", err)
		return c.Reply("Error, try again")
	}

	err = h.service.ServiceStoreToContextDB(ctx, c.Sender().ID, "user", s)
	if err != nil {
		log.Println("Error: ", err)
		return c.Reply("Error, try again")
	}
	err = h.service.ServiceStoreToContextDB(ctx, c.Sender().ID, "model", response)
	if err != nil {
		log.Println("Error: ", err)
		return c.Reply("Error, try again!")
	}
	return c.Reply(response)
}

func (h *Handler) ChatLogs(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	for _, s := range c.Message().Payload {
		if unicode.IsLetter(s) || unicode.IsSymbol(s) {
			return c.Reply("Wrong format. Use numbers")
		}
	}

	val, _ := strconv.ParseInt(c.Message().Payload, 10, 64)
	if val == 0 {
		val = 200
	}

	str, err := h.service.ServiceReadFromChatLogDB(ctx, val)
	if err != nil {
		log.Println("Error: ", err)
		return c.Reply("Error, try again later.")
	}
	response, err := h.service.ServiceGeminiGetResponseNoHistory(ctx, "(Всё что в скобках-системный промпт, твоя цель пересказать коротко что произошло в этих сообщениях, от самых левых, тоесть новых, к старым справа): "+str)
	if err != nil {
		log.Println("Error: ", err)
		return c.Reply("Error, try again later.")
	}
	return c.Reply(response + " " + c.Message().Payload)
}
