package script

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type ScriptService struct {
	Id            int
	Name          string
	ServiceTypeId int
}

var scriptServices []ScriptService

func SelectScriptServices(scriptUUID string) ([]ScriptService, error) {
	scriptServices = nil

	fmt.Println("models.selectScriptServices")
	statement, err := scrunt.Db.Prepare(
		`SELECT sst.id,
			sst.name,
			sst.service_type_id
		FROM scripts AS s
		INNER JOIN script_service_types AS sst
		ON s.id = sst.script_id
		WHERE s.uuid = $1`)
	if err != nil {
		fmt.Println("models.selectScriptServices", "Error", err)
		return nil, err
	}

	fmt.Println("models.selectScriptServices", "Execute query")
	rows, err := statement.Query(scriptUUID)
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
		err := insertScriptServices(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return scriptServices, err
	}

	fmt.Println("Final", scriptAll)

	return scriptServices, nil
}

func insertScriptServices(rows *sql.Rows) error {
	var id int
	var name string
	var serviceTypeId int

	fmt.Println("models.selectScriptsAll", "Insert Script")
	err := rows.Scan(&id, &name, &serviceTypeId)
	if err != nil {
		fmt.Println("models.selectScriptService", "Error", err)
		return err
	}

	var scriptService ScriptService

	scriptService.Id = id
	scriptService.Name = name
	scriptService.ServiceTypeId = serviceTypeId

	fmt.Println(scriptService)

	scriptServices = append(scriptServices, scriptService)

	fmt.Println(scriptAll)

	return nil
}
