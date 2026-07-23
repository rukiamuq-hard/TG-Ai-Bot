package ai

import (
	"TgAiBot/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

const url = "https://generativelanguage.googleapis.com/v1beta/models/gemini-3.1-flash-lite:generateContent?key="

type AI struct {
}

type Request struct {
	Contents []models.Content `json:"contents"`
}

type response struct {
	Candidates []struct {
		Content struct {
			Parts []models.Part `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func New() *AI {
	return &AI{}
}

func (ai *AI) GeminiGetResponse(history []models.Content, text string) (string, error) {
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

	resp, err := http.Post(url+apiKey, "application/json", bytes.NewBuffer(body))
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

func (ai *AI) GeminiGetResponseNoHistory(text string) (string, error) {
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
