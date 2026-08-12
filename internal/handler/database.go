package handler

import (
	"TgAiBot/internal/models"
	"context"
	"log"
	"strconv"
	"time"
	"unicode"

	tele "gopkg.in/telebot.v4"
)

func (h *Handler) StoreMessage(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	hist := models.History{
		Name: c.Message().Sender.FirstName + c.Message().Sender.LastName,
		UID:  c.Sender().ID,
		Text: c.Text(),
		CID:  c.Chat().ID,
		MID:  c.Message().ID,
	}

	err := h.service.ServiceStoreToChatLogDB(ctx, hist)
	if err != nil {
		log.Println("failed store to chat log db")
		return err
	}
	return nil
}

func (h *Handler) ClearMessage(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, s := range c.Message().Payload {
		if unicode.IsLetter(s) || unicode.IsSymbol(s) {
			return c.Reply("Usage: /Clear <number>")
		}
	}

	val, _ := strconv.ParseInt(c.Message().Payload, 10, 64)
	if val <= 0 {
		return c.Reply("Incorrect argument")
	}

	hist, _, err := h.service.ServiceReadFromChatLogDB(ctx, val)
	if err != nil {
		return err
	}

	var msgs []tele.Editable

	for _, ids := range hist {
		msgs = append(msgs, &tele.StoredMessage{
			ChatID:    c.Chat().ID,
			MessageID: strconv.Itoa(ids.MID),
		})
	}

	if err := h.service.ServiceDeleteFromChatLogDB(ctx, c.Chat().ID, val); err != nil {
		return c.Reply(err)
	}
	return c.Bot().DeleteMany(msgs)

}
