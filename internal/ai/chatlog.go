package ai

import (
	"TgAiBot/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

func (ai *AI) GeminiGetResponseNoHistory(ctx context.Context, text string) (string, error) { // CHATLOGS COMMAND
	apiKey := os.Getenv("AI_TOKEN")

	body, err := json.Marshal(Request{
		Contents: []models.Content{
			{Parts: []models.Part{{Text: text}}},
		},
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url+apiKey, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result response
	json.NewDecoder(resp.Body).Decode(&result)

	if len(result.Candidates) == 0 {
		return "", errors.New("Empty Gemini reponse!")
	}
	return result.Candidates[0].Content.Parts[0].Text, nil
}
