package runtime

import (
	"fmt"
	_ "modernc.org/sqlite"
)

func InsertOutput(runId int, outputType int, value string) (int64, error) {
	statement, err := db.Prepare(`INSERT INTO outputs (run_id, type, OutputValue) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(runId, outputType, value)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
