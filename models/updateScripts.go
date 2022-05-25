package models

import (
	_ "github.com/mattn/go-sqlite3"
)

func UpdateScript(id int, name string, description string, code string) (int64, error) {
	statement, err := db.Prepare(`UPDATE scripts SET name=?, description=?, code=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, description, code, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
