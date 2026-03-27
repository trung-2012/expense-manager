package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open("sqlite", "./expense.db")
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	query := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		amount REAL,
		category TEXT,
		user_id INTEGER
	);`

	_, err = DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
