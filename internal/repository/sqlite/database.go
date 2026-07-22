package database

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type SQLite struct {
	db *sql.DB
}

func New() *SQLite {
	return &SQLite{}
}

func (SQLdb *SQLite) CreateStartDB() error {
	var err error
	SQLdb.db, err = sql.Open("sqlite", "./ChatHistory.db")
	if err != nil {
		return err
	}

	SQLdb.db.SetMaxOpenConns(1)

	SQLdb.db.Exec(ContextDB)
	SQLdb.db.Exec(ChatHistoryDB)
	return nil
}

func (SQLdb *SQLite) CloseDB(stor *SQLite) {
	stor.db.Close()
}
