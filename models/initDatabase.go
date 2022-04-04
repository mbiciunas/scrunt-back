package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() error {
	database, err := sql.Open("sqlite3", ".scrunt/database/scrunt.db")
	if err != nil {
		return err
	}

	db = database

	fmt.Println(db.Ping())

	return db.Ping()
}
