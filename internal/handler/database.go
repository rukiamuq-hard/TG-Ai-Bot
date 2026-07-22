package handler

import (
	tele "gopkg.in/telebot.v4"
	"log"
)

func (h *Handler) StoreMessage(c tele.Context) error {
	err := h.service.ServiceStoreToChatLogDB(c.Sender().FirstName+" "+c.Sender().LastName, c.Text())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
