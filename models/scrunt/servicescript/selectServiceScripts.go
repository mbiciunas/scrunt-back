package servicescript

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

type ServiceScript struct {
	ScriptId          int
	ScriptName        string
	ScriptDescription string
}

var serviceScripts []ServiceScript

func SelectServiceScripts(serviceId int) ([]ServiceScript, error) {
	serviceScripts = nil

	statement, err := scrunt.Db.Prepare(
		`SELECT DISTINCT sc.id AS "script_id",
                      sc.name AS "script_name",
                      sc.description AS "script_description"
               FROM services AS s
               INNER JOIN service_types AS st ON s.service_type_id = st.id
               INNER JOIN script_service_types AS sst ON st.id = sst.service_type_id
               INNER JOIN scripts AS sc ON sst.script_id = sc.id
	           WHERE s.id = $1`)
	if err != nil {
		fmt.Println("models.selectServiceScripts", "Error", err)
		return nil, err
	}

	fmt.Println("models.selectServiceScripts", "Execute query")
	rows, err := statement.Query(serviceId)
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
		err := insertServiceScripts(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return serviceScripts, err
	}

	fmt.Println("Final", serviceScripts)

	return serviceScripts, nil
}

func insertServiceScripts(rows *sql.Rows) error {
	var scriptId int
	var scriptName string
	var scriptDescription string

	fmt.Println("models.selectServiceKeys.insertServiceKeys", "Insert ServiceKeys")
	err := rows.Scan(&scriptId, &scriptName, &scriptDescription)
	if err != nil {
		fmt.Println("models.selectServiceKeys.insertServiceKeys", "Error", err)
		return err
	}

	var serviceScript ServiceScript

	serviceScript.ScriptId = scriptId
	serviceScript.ScriptName = scriptName
	serviceScript.ScriptDescription = scriptDescription

	fmt.Println(serviceScript)

	serviceScripts = append(serviceScripts, serviceScript)

	fmt.Println(serviceScripts)

	return nil
}
