package ai

import (
	"TgAiBot/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func (ai *AI) GeminiGetResponse(ctx context.Context, history []models.Content, text string) (string, error) {
	apiKey := os.Getenv("AI_TOKEN")

	history = append(history, models.Content{
		Role:  "user",
		Parts: []models.Part{{Text: text}},
	})

	body, err := json.Marshal(Request{
		Contents: history,
	})
	if err != nil {
		return "", err
	}

	req, _ := http.NewRequestWithContext(ctx, "POST", url+apiKey, bytes.NewBuffer(body))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
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
