package script

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

type Script struct {
	Id          int
	Name        string
	Description string
	Code        string
}

func SelectScript(id int) (Script, error) {
	statement, err := store.Db.Prepare(`SELECT id, name, description, code FROM scripts WHERE id = $1`)
	if err != nil {
		fmt.Println("SelectScript - prepare", err)
		return Script{}, err
	}

	rows, err := statement.Query(id)
	if err != nil {
		fmt.Println("SelectScript - query", err)
		return Script{}, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	var script Script

	for rows.Next() {
		//var cred credential

		err := rows.Scan(&script.Id, &script.Name, &script.Description, &script.Code)
		if err != nil {
			return script, err
		}
	}
	if err = rows.Err(); err != nil {
		return script, err
	}

	return script, nil
}
