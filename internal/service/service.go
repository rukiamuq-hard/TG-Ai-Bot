package service

import (
	"TgAiBot/internal/models"
	"context"
)

type ChatLogRepository interface {
	StoreToChatLogDB(ctx context.Context, name string, text string) error
	ReadFromChatLogDB(ctx context.Context, val int64) (string, error)
}

type ContextRepository interface {
	StoreToContextDB(ctx context.Context, ser_id int64, model string, text string) error
	ReadFromContextDB(ctx context.Context, ser_id int64) ([]models.Content, error)
}

//type MongoDB interface {}

type AI interface {
	GeminiGetResponse(ctx context.Context, istory []models.Content, text string) (string, error)
	GeminiGetResponseNoHistory(ctx context.Context, ext string) (string, error)
}

type Service struct {
	DBChatLog ChatLogRepository
	DBContext ContextRepository
	//	LogsDB MongoDB
	ai AI
}

func New(DBChatLog ChatLogRepository, DBContext ContextRepository /*LogsDB MongoDB*/, ai AI) *Service {
	return &Service{
		DBChatLog: DBChatLog,
		DBContext: DBContext,
		//		LogsDB: LogsDB,
		ai: ai,
	}
}
