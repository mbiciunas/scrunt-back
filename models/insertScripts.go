package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InsertScript(name string, description string, code string) (int64, error) {
	statement, err := db.Prepare(`INSERT INTO scripts (name, description, code) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, description, code)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
