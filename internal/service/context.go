package service

import (
	"TgAiBot/internal/models"
)

func (svc *Service) ServiceStoreToContextDB(user_id int64, model string, text string) error {
	return svc.sqlDB.StoreToContextDB(user_id, model, text)
}

func (svc *Service) ServiceReadFromContextDB(user_id int64) ([]models.Content, error) {
	return svc.sqlDB.ReadFromContextDB(user_id)
}
