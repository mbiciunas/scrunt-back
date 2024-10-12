package server

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

func DeleteServer(id int) (int64, error) {
	tx, err := store.Db.Begin()
	if err != nil {
		return 0, err
	}

	err = deleteServerCredentialByServer(tx, id)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return 0, rollbackErr
		}
		return 0, err
	}

	rows, err := deleteServer(tx, id)
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

func deleteServer(tx *sql.Tx, id int) (int64, error) {
	statement, err := tx.Prepare(`DELETE FROM servers WHERE id = ?`)
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

func deleteServerCredentialByServer(tx *sql.Tx, id int) error {
	statement, err := tx.Prepare(`DELETE FROM server_credentials WHERE server_id = ?`)
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
