package service

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models"
)

func InsertService(name string, description string, address string, port uint, serviceTypeId uint) (int64, error) {
	statement, err := models.Db.Prepare(`INSERT INTO services (name, description, address, port, service_type_id) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, description, address, port, serviceTypeId)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
