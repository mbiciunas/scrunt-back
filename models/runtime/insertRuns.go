package runtime

import (
	"fmt"
	_ "modernc.org/sqlite"
)

func InsertRuns(scriptId int, status int) (int64, error) {
	statement, err := db.Prepare(`INSERT INTO runs (script_id, status) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(scriptId, status)
	if err != nil {
		return 0, err
	} else {
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("InsertRuns - Insert id: ", lastInsertId)

		return result.LastInsertId()
	}
}
