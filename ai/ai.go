package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type request struct {
	Contents []content `json:"contents"`
}

type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text"`
}

type response struct {
	Candidates []struct {
		Content struct {
			Parts []part `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func init() {
	if err := godotenv.Load("../TokensChatId.env"); err != nil {
		log.Print("Error load .env")
	}
}

func GeminiGetResponse(text string) (string, error) {
	apiKey := os.Getenv("AI_TOKEN")
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/gemini-3.1-flash-lite:generateContent?key=%s", apiKey)
	body, err := json.Marshal(request{
		Contents: []content{
			{Parts: []part{{Text: text}}},
		},
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println(err)
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
