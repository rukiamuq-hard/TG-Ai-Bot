package handler

import (
	"TgAiBot/internal/models"
	"context"
)

type Service interface {
	ServiceStoreToChatLogDB(ctx context.Context, name string, text string) error
	ServiceReadFromChatLogDB(ctx context.Context, val int64) (string, error)
	ServiceStoreToContextDB(ctx context.Context, user_id int64, model string, text string) error
	ServiceReadFromContextDB(ctx context.Context, user_id int64) ([]models.Content, error)
	ServiceGeminiGetResponse(ctx context.Context, history []models.Content, text string) (string, error)
	ServiceGeminiGetResponseNoHistory(ctx context.Context, text string) (string, error)
}

type Handler struct {
	service Service
}

func New(svc Service) *Handler {
	return &Handler{service: svc}
}
