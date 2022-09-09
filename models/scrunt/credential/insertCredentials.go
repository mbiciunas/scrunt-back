package credential

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models"
)

func InsertCredential(name string, credtype string, credential string) (int64, error) {
	statement, err := models.Db.Prepare(`INSERT INTO credentials (name, type, credential) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, credtype, credential)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
