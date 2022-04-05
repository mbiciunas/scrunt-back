package models

import (
	_ "github.com/mattn/go-sqlite3"
)

type Script struct {
	Id     int
	Name   string
	Script string
}

func SelectScript(id int) (Script, error) {
	statement, err := db.Prepare(`SELECT id, name, script FROM scripts WHERE id = $1`)
	if err != nil {
		return Script{}, err
	}

	rows, err := statement.Query(id)
	if err != nil {
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

		err := rows.Scan(&script.Id, &script.Name, &script.Script)
		if err != nil {
			return script, err
		}
	}
	if err = rows.Err(); err != nil {
		return script, err
	}

	return script, nil
}
