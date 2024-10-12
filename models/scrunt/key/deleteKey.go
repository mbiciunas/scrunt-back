package key

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

func DeleteKey(id int) (int64, error) {
	tx, err := scrunt.Db.Begin()
	if err != nil {
		return 0, err
	}

	err = deleteServiceKeyByKey(tx, id)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	rows, err := deleteKey(tx, id)
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

func deleteKey(tx *sql.Tx, id int) (int64, error) {
	statement, err := tx.Prepare(`DELETE FROM keys WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(id)
	if err != nil {
		return 0, err
	} else {
		fmt.Print(result.RowsAffected())
		return result.RowsAffected()
	}
}

func deleteServiceKeyByKey(tx *sql.Tx, id int) error {
	statement, err := tx.Prepare(`DELETE FROM service_keys WHERE key_id = ?`)
	if err != nil {
		return err
	}

	_, err = statement.Exec(id)
	if err != nil {
		return err
	} else {
		return nil
	}
}
