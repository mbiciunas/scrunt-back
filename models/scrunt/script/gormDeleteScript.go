package script

import (
	"errors"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
	"strings"
)

func GormDeleteScript(tx *gorm.DB, id int) (int64, error) {
	rows, err := deleteScript1(tx, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	} else {
		if rows < 0 {
			tx.Rollback()
			return rows, errors.New("no rows deleted")
		}
	}

	return rows, nil
}

func deleteScript1(tx *gorm.DB, id int) (int64, error) {
	var query strings.Builder

	query.WriteString("DELETE FROM scripts ")
	query.WriteString("WHERE id = ? ")

	errGorm := tx.Exec(query.String(), id)

	if errGorm.Error != nil {
		return 0, errGorm.Error
	}

	return errGorm.RowsAffected, nil
}
