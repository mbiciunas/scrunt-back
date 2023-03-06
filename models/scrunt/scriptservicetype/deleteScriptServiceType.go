package scriptservicetype

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models"
)

func DeleteScriptServiceType(serviceKeyId int) (int64, error) {
	tx, err := models.Db.Begin()
	if err != nil {
		return 0, err
	}

	rows, err := deleteScriptServiceType(tx, serviceKeyId)
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

func deleteScriptServiceType(tx *sql.Tx, id int) (int64, error) {
	statement, err := tx.Prepare(`DELETE FROM script_service_types WHERE id = ?`)
	fmt.Println("deleteScriptServiceType.deleteScriptServiceType - statement", statement, id)
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
