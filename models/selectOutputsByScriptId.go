package models

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Output struct {
	Id              int
	ScriptId        int
	DatetimeCreated int
	Output          string
	Log             string
}

func SelectOutputByScriptId(scriptId int) (Output, error) {
	statement, err := db.Prepare(`SELECT id, script_id, datetime_created, output, log FROM outputs WHERE script_id = $1`)
	if err != nil {
		fmt.Println("SelectOutputByScriptId - prepare", err)
		return Output{}, err
	}

	rows, err := statement.Query(scriptId)
	if err != nil {
		fmt.Println("SelectOutputByScriptId - query", err)
		return Output{}, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	var output Output

	for rows.Next() {
		err := rows.Scan(&output.Id, &output.ScriptId, &output.DatetimeCreated, &output.Output, &output.Log)
		if err != nil {
			return output, err
		}
	}
	if err = rows.Err(); err != nil {
		return output, err
	}

	return output, nil
}
