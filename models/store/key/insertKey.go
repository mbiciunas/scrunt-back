package key

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

func InsertKey(name string, description string, keytype uint, keyprivate string, keypublic string) (int64, error) {
	statement, err := store.Db.Prepare(`INSERT INTO keys (name, description, key_type_id, private, public) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, description, keytype, keyprivate, keypublic)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
