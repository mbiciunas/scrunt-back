package models

import (
	_ "github.com/mattn/go-sqlite3"
)

func DeleteServerCredential(id int) (int64, error) {
	statement, err := db.Prepare(`DELETE FROM server_credentials WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
