package app

import (
	"TgAiBot/internal/ai"
	"TgAiBot/internal/handler"
	"TgAiBot/internal/repository/chatlogdb"
	"TgAiBot/internal/repository/contextdb"
	"TgAiBot/internal/service"

	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

type App struct {
	svc    *service.Service
	ctxDB  *database.SQLite
	logsDB *mongodb.LogsDB
	ai     *ai.AI
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No env found, using standart env!")
	}
}

func New() *App {
	return &App{}
}

func (app *App) Start() error {
	token := os.Getenv("TOKEN")

	app.ctxDB = database.New()
	app.ai = ai.New()
	app.logsDB = mongodb.New()

	_Service := service.New(app.logsDB, app.ctxDB, app.ai) // first arg is ChatLogDB, second is ContextDB
	h := handler.New(_Service)

	err := app.logsDB.ConnectDB()
	if err != nil {
		return err
	}

	err = app.ctxDB.CreateStartDB()
	if err != nil {
		return err
	}

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 5 * time.Second}, // how fast bot check message
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return err
	}

	b.Handle(tele.OnText, h.StoreMessage)

	b.Handle("/Gemini", h.GeminiGetResp)

	b.Handle("/ChatLogs", h.GetHistory)

	b.Handle("/Clear", h.ClearMessage)

	b.Start()

	return nil
}

func (app *App) Close() {
	app.ctxDB.CloseDB()
	app.logsDB.Close()
}
