package database

import (
	"TgAiBot/internal/models"
	"fmt"
	_ "modernc.org/sqlite"
)

const ContextDB = `
	CREATE TABLE IF NOT EXISTS Context(
		id INTEGER PRIMARY KEY,
		role TEXT,
		user_id BIGINT,
		text TEXT
	);`

const InsertContextByIDRT = `INSERT INTO Context (user_id, role, text) VALUES (?, ?, ?)`

const ReadFromContextDB = `SELECT text, role FROM Context WHERE user_id = ? ORDER BY id ASC LIMIT 20`

func (SQLdb *SQLite) StoreToContextDB(user_id int64, model string, text string) error {
	_, err := SQLdb.db.Exec(InsertContextByIDRT, user_id, model, text)
	if err != nil {
		return err
	}
	fmt.Println("USER_ID: ", user_id, "\n MODEL: ", model, "\n", " TEXT:", text)
	return nil
}

func (SQLdb *SQLite) ReadFromContextDB(user_id int64) ([]models.Content, error) {
	rows, err := SQLdb.db.Query(ReadFromContextDB, user_id)
	if err != nil {
		return nil, err
	}

	var history []models.Content
	for rows.Next() {
		var r, m string
		rows.Scan(&m, &r)
		history = append(history, models.Content{
			Role:  r,
			Parts: []models.Part{{Text: m}},
		})
	}

	return history, rows.Err()
}
