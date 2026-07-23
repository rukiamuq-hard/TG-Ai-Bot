package service

import (
	"TgAiBot/internal/models"
)

func (svc *Service) ServiceGeminiGetResponse(history []models.Content, text string) (string, error) {
	return svc.ai.GeminiGetResponse(history, text)
}

func (svc *Service) ServiceGeminiGetResponseNoHistory(text string) (string, error) {
	return svc.ai.GeminiGetResponseNoHistory(text)
}
