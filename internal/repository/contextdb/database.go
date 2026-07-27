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
	return nil
}

func (SQLdb *SQLite) CloseDB() {
	SQLdb.db.Close()
}
