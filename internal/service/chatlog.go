package service

import (
	"TgAiBot/internal/models"
	"context"
)

func (svc *Service) ServiceStoreToChatLogDB(ctx context.Context, hist models.History) error {
	return svc.DBChatLog.StoreToChatLogDB(ctx, hist)
}

func (svc *Service) ServiceReadFromChatLogDB(ctx context.Context, val int64) ([]models.History, string, error) {
	return svc.DBChatLog.ReadFromChatLogDB(ctx, val)
}

func (svc *Service) ServiceDeleteFromChatLogDB(ctx context.Context, chat_id int64, val int64) error {
	return svc.DBChatLog.DeleteMessage(ctx, chat_id, val)
}
