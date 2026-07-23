package service

import (
	"TgAiBot/internal/models"
	"context"
	"log"
	"os"
)

func (svc *Service) ServiceGeminiGetResponse(ctx context.Context, history []models.Content, text string) (string, error) {
	prompt, err := os.ReadFile("prompt.txt")
	if err != nil {
		log.Println("Default working, without prompt", err)
	}
	return svc.ai.GeminiGetResponse(ctx, history, "[PROMPT]:"+string(prompt)+text)
}

func (svc *Service) ServiceGeminiGetResponseNoHistory(ctx context.Context, text string) (string, error) {
	return svc.ai.GeminiGetResponseNoHistory(ctx, text)
}
