package script

import (
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

func UpdateScript(id int, name string, description string, code string) (int64, error) {
	statement, err := store.Db.Prepare(`UPDATE scripts SET name=?, description=?, code=? WHERE id = ?`)
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
