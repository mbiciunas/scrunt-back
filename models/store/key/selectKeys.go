package key

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

type Key struct {
	Id          int
	Name        string
	Description string
	Type        int
	KeyPublic   string
	KeyPrivate  string
}

func SelectKey(id int) (Key, error) {
	statement, err := store.Db.Prepare(`SELECT id, name, description, key_type_id, public, private FROM keys WHERE id = $1`)
	if err != nil {
		fmt.Println("SelectKey - prepare", err)
		return Key{}, err
	}

	rows, err := statement.Query(id)
	if err != nil {
		fmt.Println("SelectKey - query", err)
		return Key{}, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	var key Key

	for rows.Next() {
		//var cred credential

		err := rows.Scan(&key.Id, &key.Name, &key.Description, &key.Type, &key.KeyPublic, &key.KeyPrivate)
		if err != nil {
			return key, err
		}
	}
	if err = rows.Err(); err != nil {
		return key, err
	}

	return key, nil
}
