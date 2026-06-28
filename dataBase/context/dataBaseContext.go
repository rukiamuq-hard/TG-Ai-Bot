package dataBaseContext

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "modernc.org/sqlite"
)

type Storage struct {
	db *sql.DB
}

func CreateStartDB() (*Storage, error) {
	db, err := sql.Open("sqlite", "./ChatHistory.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ContextForAI := `
	CREATE TABLE IF NOT EXISTS contextAI(
		user_id BIGINT,
		text TEXT
	);`

	ChatHistoryDB := `
	CREATE TABLE IF NOT EXISTS ChatLogs(
		name TEXT,
		text TEXT
	);`

	db.Exec(ContextForAI)
	db.Exec(ChatHistoryDB)
	return &Storage{db: db}, nil

}
func StoreToChatLogDB(stor *Storage, name string, text string) {
	query := "INSERT INTO ChatLogs(name, text) VALUES (?, ?)"
	_, err := stor.db.Exec(query, name, text)
	if err != nil {
		log.Fatal("Can`t write to db: ", err)
	}
}

func ReadFromChatLogDB(stor *Storage) string {
	query := "SELECT name, text FROM ChatLogs ORDER BY rowid ASC LIMIT 200"
	rows, err := stor.db.Query(query)
	if err != nil {
		log.Fatal("Can`t read DB: ", err)
		return ""
	}
	var s strings.Builder
	for rows.Next() {
		var n string
		var t string
		rows.Scan(&n, &t)
		s.WriteString("NAME: " + n + " TEXT: " + t + "\n")
	}
	return s.String()
}

func StoreToContextDB(stor *Storage, user_id int64, text string) {
	InsertQuery := `INSERT INTO contextAI (user_id, text) VALUES (?, ?)`
	_, err := stor.db.Exec(InsertQuery, user_id, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("USER_ID: ", user_id, "\n", "TEXT:", text)
}

func ReadFromContextDB(stor *Storage, user_id int64) string {
	ReadQuery := `SELECT text FROM contextAI WHERE user_id = ? ORDER BY rowid ASC LIMIT 20`
	rows, err := stor.db.Query(ReadQuery, user_id)
	if err != nil {
		log.Fatal(err)
	}

	var s strings.Builder
	for rows.Next() {
		var r string
		rows.Scan(&r)
		s.WriteString(r + " ")
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return s.String()
}

func CloseDB(stor *Storage) {
	stor.db.Close()
}
