package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type request struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Role  string `json:"role"`
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type response struct {
	Candidates []struct {
		Content struct {
			Parts []Part `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func init() {
	if err := godotenv.Load("../TokensChatId.env"); err != nil {
		fmt.Println("error", err)
	}
}

func GeminiGetResponse(history []Content, text string) (string, error) {
	apiKey := os.Getenv("AI_TOKEN")
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-3.1-flash-lite:generateContent?key=%s", apiKey)

	history = append(history, Content{
		Role:  "user",
		Parts: []Part{{Text: text}},
	})

	body, err := json.Marshal(request{
		Contents: history,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
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

func GeminiGetResponseNoHistory(text string) (string, error) {
	apiKey := os.Getenv("AI_TOKEN")
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-3.1-flash-lite:generateContent?key=%s", apiKey)

	body, err := json.Marshal(request{
		Contents: []Content{
			{Parts: []Part{{Text: text}}},
		},
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
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
