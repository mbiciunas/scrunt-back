package script

import (
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"strings"
)

func GormDeleteScript(tx *gorm.DB, id int) (int64, error) {
	//tx := scrunt.GormDB.Begin()
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//if err := tx.Error; err != nil {
	//	fmt.Println("models.scrunt.script.gormDeleteScript", "transaction error", err, tx.Error)
	//	return 0, err
	//}

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

	//tx.Commit()
	//if tx.Error != nil {
	//	return 0, tx.Error
	//}

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
