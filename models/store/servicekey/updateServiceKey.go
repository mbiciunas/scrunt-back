package servicekey

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

func UpdateServiceKey(serviceKeyId int, serviceId int, keyId int) (int64, error) {
	statement, err := store.Db.Prepare(`UPDATE service_keys SET key_id=?, service_id=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(keyId, serviceId, serviceKeyId)
	if err != nil {
		fmt.Println("updateServiceKey.UpdateServiceKey - Error:", err)
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
