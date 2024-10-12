package key

import (
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

func UpdateKey(id int, name string, description string, keyType int64, keyPublic string, keyPrivate string) (int64, error) {
	statement, err := store.Db.Prepare(`UPDATE keys SET name=?, description=?, key_type_id=?, public=?, private=? WHERE id = ?`)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(name, description, keyType, keyPublic, keyPrivate, id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
