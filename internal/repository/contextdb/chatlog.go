package database

/*
import (
	"context"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

const ChatHistoryDB = `
	CREATE TABLE IF NOT EXISTS ChatLogs(
		id INTEGER PRIMARY KEY,
		name TEXT,
		text TEXT
	);`

const InsertChatLogByNT = "INSERT INTO ChatLogs(name, text) VALUES (?, ?)"

const ReadChatLogByVAL = "SELECT name, text FROM ChatLogs ORDER BY id DESC LIMIT ?"

func (SQLdb *SQLite) StoreToChatLogDB(ctx context.Context, name string, text string) error {
	_, err := SQLdb.db.ExecContext(ctx, InsertChatLogByNT, name, text)
	fmt.Println("CHATLOG: ", name, " TEXT: ", text)
	return err
}

func (SQLdb *SQLite) ReadFromChatLogDB(ctx context.Context, val int64) (string, error) {
	if val == 0 {
		val = 200
	}
	rows, err := SQLdb.db.QueryContext(ctx, ReadChatLogByVAL, val)
	if err != nil {
		return "", err
	}
	var s strings.Builder
	for rows.Next() {
		var n string
		var t string
		rows.Scan(&n, &t)
		s.WriteString("NAME: ")
		s.WriteString(n)
		s.WriteString(" TEXT: ")
		s.WriteString(t)
		s.WriteString("\n")
	}
	return s.String(), nil
}
*/
