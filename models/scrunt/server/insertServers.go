package server

import (
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func InsertServer(name string, address string) (int64, error) {
	statement, err := scrunt.Db.Prepare(`INSERT INTO servers (name, address) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, address)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
