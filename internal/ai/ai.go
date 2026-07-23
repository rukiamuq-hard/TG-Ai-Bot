package ai

import (
	"TgAiBot/internal/models"
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
