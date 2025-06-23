package servicekey

import (
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func InsertServiceKey(keyId uint, serviceId uint) (int64, error) {
	fmt.Println("insertServiceKey.InsertServiceKey - keyId:", keyId, "serviceId", serviceId)
	statement, err := scrunt.Db.Prepare(`INSERT INTO service_keys (key_id, service_id) VALUES (?, ?)`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(keyId, serviceId)
	if err != nil {
		return 0, err
	} else {
		rowsAffected, _ := result.RowsAffected()
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Result: ", rowsAffected, lastInsertId)

		return result.LastInsertId()
	}
}
