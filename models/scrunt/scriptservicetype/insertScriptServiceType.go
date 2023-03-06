package scriptservicetype

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models"
)

func InsertScriptServiceType(scriptId uint, serviceTypeId uint, name string) (int64, error) {
	fmt.Println("insertScriptServiceType.InsertScriptServiceType - scriptId:", scriptId, "serviceTypeId", serviceTypeId, "name", name)
	statement, err := models.Db.Prepare(`INSERT INTO script_service_types (script_id, service_type_id, name) VALUES (?, ?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(scriptId, serviceTypeId, name)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
