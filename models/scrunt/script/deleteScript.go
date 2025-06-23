package script

import (
	"database/sql"
	"errors"
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func DeleteScript(id int) (int64, error) {
	tx, err := scrunt.Db.Begin()
	if err != nil {
		return 0, err
	}

	rows, err := deleteScript(tx, id)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rows, rollbackErr
		}
		return 0, err
	} else {
		if rows == 0 {
			return rows, errors.New("no rows deleted")
		}
	}

	err = tx.Commit()
	if err != nil {
		return rows, err
	}

	return rows, nil
}

func deleteScript(tx *sql.Tx, id int) (int64, error) {
	statement, err := tx.Prepare(`DELETE FROM scripts WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(id)
	if err != nil {
		return 0, err
	} else {
		fmt.Println(result.RowsAffected())
		return result.RowsAffected()
	}
}
