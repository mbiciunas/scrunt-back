package script

import (
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func UpdateScript(id int, name string, description string, code string) (int64, error) {
	statement, err := scrunt.Db.Prepare(`UPDATE scripts SET name=?, description=?, code=? WHERE id = ?`)
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
