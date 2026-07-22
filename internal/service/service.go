package service

import (
	"TgAiBot/internal/models"
)

type SQLite interface {
	StoreToChatLogDB(name string, text string) error
	ReadFromChatLogDB(val int64) (string, error)
	StoreToContextDB(user_id int64, model string, text string) error
	ReadFromContextDB(user_id int64) ([]models.Content, error)
}

//type MongoDB interface {}

type Service struct {
	sqlDB SQLite
	//	LogsDB MongoDB
}

func New(sqlDB SQLite /*LogsDB MongoDB*/) *Service {
	return &Service{
		sqlDB: sqlDB,
		//		LogsDB: LogsDB,
	}
}
