package handler

import (
	"TgAiBot/internal/models"
)

type Service interface {
	ServiceStoreToChatLogDB(name string, text string) error
	ServiceReadFromChatLogDB(val int64) (string, error)
	ServiceStoreToContextDB(user_id int64, model string, text string) error
	ServiceReadFromContextDB(user_id int64) ([]models.Content, error)
}

type Handler struct {
	service Service
}

func New(svc Service) *Handler {
	return &Handler{service: svc}
}
