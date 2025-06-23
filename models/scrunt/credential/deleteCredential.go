package credential

import (
	"database/sql"
	"errors"
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func DeleteCredential(id int) (int64, error) {
	tx, err := scrunt.Db.Begin()
	if err != nil {
		return 0, err
	}

	err = deleteServerCredentialByCredential(tx, id)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	rows, err := deleteCredential(tx, id)
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

func deleteCredential(tx *sql.Tx, id int) (int64, error) {
	statement, err := tx.Prepare(`DELETE FROM credentials WHERE id = ?`)
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

func deleteServerCredentialByCredential(tx *sql.Tx, id int) error {
	statement, err := tx.Prepare(`DELETE FROM server_credentials WHERE credential_id = ?`)
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
