package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type ScriptAll struct {
	Id     int
	Name   string
	Script string
}

var scriptAll []ScriptAll

func SelectScriptsAll() ([]ScriptAll, error) {
	scriptAll = nil

	fmt.Println("models.selectScriptsAll")
	statement, err := db.Prepare(
		`SELECT s.id,
			s.name,
			s.script
		FROM scripts s`)
	if err != nil {
		return nil, err
	}

	fmt.Println("models.selectScriptsAll", "Execute query")
	rows, err := statement.Query()
	if err != nil {
		fmt.Println("Query error", err)
		return nil, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	for rows.Next() {
		err := insertScript(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return scriptAll, err
	}

	fmt.Println("Final", credentialAll)

	return scriptAll, nil
}

func insertScript(rows *sql.Rows) error {
	var id int
	var name string
	var script string

	fmt.Println("models.selectScriptsAll", "Insert Script")
	err := rows.Scan(&id, &name, &script)
	if err != nil {
		return err
	}

	var scriptRow ScriptAll

	scriptRow.Id = id
	scriptRow.Name = name
	scriptRow.Script = script

	fmt.Println(scriptRow)

	scriptAll = append(scriptAll, scriptRow)

	fmt.Println(credentialAll)

	return nil
}
