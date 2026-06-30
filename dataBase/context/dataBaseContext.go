package dataBaseContext

import (
	"TgAiBot/ai"
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"strings"
)

type Storage struct {
	db *sql.DB
}

func CreateStartDB() (*Storage, error) {
	db, err := sql.Open("sqlite", "./ChatHistory.db")
	if err != nil {
		return nil, err
	}

	ContextForAI := `
	CREATE TABLE IF NOT EXISTS contextAI(
		id INTEGER PRIMARY KEY,
		role TEXT,
		user_id BIGINT,
		text TEXT
	);`

	ChatHistoryDB := `
	CREATE TABLE IF NOT EXISTS ChatLogs(
		id INTEGER PRIMARY KEY,
		name TEXT,
		text TEXT
	);`

	db.Exec(ContextForAI)
	db.Exec(ChatHistoryDB)
	return &Storage{db: db}, nil

}
func StoreToChatLogDB(stor *Storage, name string, text string) error {
	query := "INSERT INTO ChatLogs(name, text) VALUES (?, ?)"
	_, err := stor.db.Exec(query, name, text)
	fmt.Println("CHATLOG: ", name, " TEXT: ", text)
	return err
}

func ReadFromChatLogDB(stor *Storage, val int64) (string, error) {
	if val == 0 {
		val = 200
	}
	rows, err := stor.db.Query("SELECT name, text FROM ChatLogs ORDER BY id DESC LIMIT ?", val)
	if err != nil {
		return "", err
	}
	var s strings.Builder
	for rows.Next() {
		var n string
		var t string
		rows.Scan(&n, &t)
		s.WriteString("NAME: " + n + " TEXT: " + t + "\n")
	}
	return s.String(), nil
}

func StoreToContextDB(stor *Storage, user_id int64, model string, text string) error {
	InsertQuery := `INSERT INTO contextAI (user_id, role, text) VALUES (?, ?, ?)`
	_, err := stor.db.Exec(InsertQuery, user_id, model, text)
	if err != nil {
		return err
	}
	fmt.Println("USER_ID: ", user_id, "\n MODEL: ", model, "\n", " TEXT:", text)
	return nil
}

func ReadFromContextDB(stor *Storage, user_id int64) ([]ai.Content, error) {
	ReadQuery := `SELECT text, role FROM contextAI WHERE user_id = ? ORDER BY id ASC LIMIT 20`
	rows, err := stor.db.Query(ReadQuery, user_id)
	if err != nil {
		return nil, err
	}

	var history []ai.Content
	for rows.Next() {
		var r, m string
		rows.Scan(&m, &r)
		history = append(history, ai.Content{
			Role:  r,
			Parts: []ai.Part{{Text: m}},
		})
	}

	return history, rows.Err()
}

func CloseDB(stor *Storage) {
	stor.db.Close()
}
