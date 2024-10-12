package service

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

func UpdateService(id int, name string, description string, address string, port int, serviceTypeId int, cloudId int) (int64, error) {
	statement, err := scrunt.Db.Prepare(`UPDATE services SET name=?, description=?, address=?, port=?, service_type_id=?, cloud_id=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, description, address, port, serviceTypeId, cloudId, id)
	if err != nil {
		fmt.Println("updateServices.UpdateService - Error:", err)
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
