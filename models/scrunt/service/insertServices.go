package service

import (
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func InsertService(name string, description string, address string, port uint, serviceTypeId uint) (int64, error) {
	statement, err := scrunt.Db.Prepare(`INSERT INTO services (name, description, address, port, service_type_id) VALUES (?, ?, ?, ?, ?)`)
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
