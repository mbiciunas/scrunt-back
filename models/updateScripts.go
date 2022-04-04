package models

import (
	_ "github.com/mattn/go-sqlite3"
)

func UpdateScript(id int, name string, script string) (int64, error) {
	statement, err := db.Prepare(`UPDATE scripts SET name=?, script=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, script, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
