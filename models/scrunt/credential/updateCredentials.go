package credential

import (
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

func UpdateCredential(id int, name string, credential string) (int64, error) {
	statement, err := scrunt.Db.Prepare(`UPDATE credentials SET name=?, credential=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, credential, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
