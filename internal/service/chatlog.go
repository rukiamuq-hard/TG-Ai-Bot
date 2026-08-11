package service

import (
	"TgAiBot/internal/models"
	"context"
)

func (svc *Service) ServiceStoreToChatLogDB(ctx context.Context, name string, text string, MID string, CID string) error {
	hist := models.History{name, text, MID, CID}
	return svc.DBChatLog.StoreToChatLogDB(ctx, hist)
}

func (svc *Service) ServiceReadFromChatLogDB(ctx context.Context, val int64) ([]models.History, string, error) {
	return svc.DBChatLog.ReadFromChatLogDB(ctx, val)
}

func (svc *Service) ServiceDeleteFromChatLogDB(ctx context.Context, val int64) error {
	return svc.DBChatLog.DeleteMessage(ctx, val)
}
