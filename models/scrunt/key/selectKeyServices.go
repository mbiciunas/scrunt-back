package key

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type KeyService struct {
	ServiceKeyId    int
	ServiceId       int
	Name            string
	Description     string
	ServiceTypeId   int
	ServiceTypeName string
}

var keyServices []KeyService

func SelectKeyServices(keyId int) ([]KeyService, error) {
	keyServices = nil

	fmt.Println("models.selectKeyServices")
	statement, err := scrunt.Db.Prepare(
		`SELECT sk.id AS "service_key_id",
			s.id AS "key_id",
			s.name,
			s.description,
			s.service_type_id,
			st.name AS "service_type_name"
		FROM service_keys sk
		INNER JOIN services s ON s.id = sk.service_id
		INNER JOIN service_types st ON st.id = s.service_type_id
		WHERE sk.key_id = $1`)
	if err != nil {
		fmt.Println("models.selectKeyServices", "Error", err)
		return nil, err
	}

	fmt.Println("models.selectKeyServices", "Execute query")
	rows, err := statement.Query(keyId)
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
		err := insertKeyServices(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return keyServices, err
	}

	fmt.Println("Final", keyServices)

	return keyServices, nil
}

func insertKeyServices(rows *sql.Rows) error {
	var serviceKeyId int
	var serviceId int
	var name string
	var description string
	var serviceTypeId int
	var serviceTypeName string

	fmt.Println("models.selectKeyServices.insertKeyServices", "Insert ServiceKeys")
	err := rows.Scan(&serviceKeyId, &serviceId, &name, &description, &serviceTypeId, &serviceTypeName)
	if err != nil {
		fmt.Println("models.selectServiceKeys.insertServiceKeys", "Error", err)
		return err
	}

	var keyService KeyService

	keyService.ServiceKeyId = serviceKeyId
	keyService.ServiceId = serviceId
	keyService.Name = name
	keyService.Description = description
	keyService.ServiceTypeId = serviceTypeId
	keyService.ServiceTypeName = serviceTypeName

	fmt.Println(keyService)

	keyServices = append(keyServices, keyService)

	fmt.Println(keyServices)

	return nil
}
