package server

import (
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

func UpdateServer(id int, name string, address string) (int64, error) {
	statement, err := store.Db.Prepare(`UPDATE servers SET name=?, address=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, address, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
