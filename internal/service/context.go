package service

import (
	"TgAiBot/internal/models"
	"context"
)

func (svc *Service) ServiceStoreToContextDB(ctx context.Context, user_id int64, model string, text string) error {
	return svc.DBContext.StoreToContextDB(ctx, user_id, model, text)
}

func (svc *Service) ServiceReadFromContextDB(ctx context.Context, user_id int64) ([]models.Content, error) {
	return svc.DBContext.ReadFromContextDB(ctx, user_id)
}
