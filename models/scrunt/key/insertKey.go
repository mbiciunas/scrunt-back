package key

import (
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func InsertKey(name string, description string, keyType uint, keyPrivate string, keyPublic string) (int64, error) {
	statement, err := scrunt.Db.Prepare(`INSERT INTO keys (name, description, key_type_id, private, public) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, description, keyType, keyPrivate, keyPublic)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
