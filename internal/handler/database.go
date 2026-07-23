package handler

import (
	"context"
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

func (h *Handler) StoreMessage(c tele.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := h.service.ServiceStoreToChatLogDB(ctx, c.Sender().FirstName+" "+c.Sender().LastName, c.Text())
	if err != nil {
		log.Println("failed store to chat log db")
		return err
	}
	return nil
}
