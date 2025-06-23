package servercredential

import (
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func DeleteServerCredential(id int) (int64, error) {
	statement, err := scrunt.Db.Prepare(`DELETE FROM server_credentials WHERE id = ?`)
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
