package app

import (
	"TgAiBot/internal/handler"
	"TgAiBot/internal/repository/sqlite"
	"TgAiBot/internal/service"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

type App struct {
	svc   *service.Service
	SQLdb *database.SQLite
	//	LogsDB //mongodb
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error load env: ", err)
	}
}

func New() *App {
	return &App{}
}

func (app *App) Start() error {
	app = &App{}
	token := os.Getenv("TOKEN")

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 5 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	prompt, err := os.ReadFile("prompt.txt")
	if err != nil {
		log.Println("Default working, without prompt", err)
	}

	app.SQLdb = database.New()
	_Service := service.New(app.SQLdb)
	h := handler.New(_Service) // need service

	b.Handle("/Gemini", h.GeminiGetResp)

	b.Handle(tele.OnText, h.StoreMessage)

	b.Handle("/ChatLogs", h.ChatLogs)

	b.Start()
}
