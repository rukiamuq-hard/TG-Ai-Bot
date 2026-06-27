package database

import (
	"database/sql"
	"fmt"
	"log"

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

	Table := `
	CREATE TABLE IF NOT EXISTS contextAI(
		user_id BIGINT,
		text TEXT
	);`

	_, err = db.Exec(Table)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Storage{db: db}, nil
}

func InsertToDB(stor *Storage, user_id int64, text string) {
	InsertQuery := `INSERT INTO contextAI (user_id, text) VALUES (?, ?)`
	_, err := stor.db.Exec(InsertQuery, user_id, text)
	if err != nil {
		log.Fatal(err)
	}

}

func ReadFromDB(stor *Storage, user_id int64) string {
	ReadQuery := `SELECT text FROM contextAI WHERE user_id = ? ORDER BY rowid DESC LIMIT 20`
	rows, err := stor.db.Query(ReadQuery, user_id)
	if err != nil {
		log.Fatal(err)
	}

	var s string

	for rows.Next() {
		var r string
		err := rows.Scan(&r)
		if err != nil {
			log.Fatal(err)
		}
		s += r + " "
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	return s
}

func CloseDB(stor *Storage) {
	stor.db.Close()
}
