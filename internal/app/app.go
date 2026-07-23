package app

import (
	"TgAiBot/internal/ai"
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
	//	LogsDB mongodb
	ai *ai.AI
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
	token := os.Getenv("TOKEN")

	app.SQLdb = database.New()
	app.ai = ai.New()
	_Service := service.New(app.SQLdb, app.ai)
	h := handler.New(_Service) // need service

	err := app.SQLdb.CreateStartDB()
	if err != nil {
		return err
	}
	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 5 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	b.Handle("/Gemini", h.GeminiGetResp)

	b.Handle(tele.OnText, h.StoreMessage)

	b.Handle("/ChatLogs", h.ChatLogs)

	b.Start()

	return nil
}

func (app *App) Close() {
	app.SQLdb.CloseDB()
}
