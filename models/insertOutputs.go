package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InsertOutput(scriptId int, output string, log string) (int64, error) {
	statement, err := db.Prepare(`INSERT INTO outputs (script_id, output, log) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(scriptId, output, log)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
