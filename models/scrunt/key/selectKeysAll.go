package key

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models"
)

type service struct {
	Id   int
	Name string
}

type KeyAll struct {
	Id          int
	Name        string
	Description string
	KeyTypeId   int
	KeyTypeName string
	Services    []service
}

var keyAll []KeyAll
var oldKeyId int

func SelectKeysAll() ([]KeyAll, error) {
	keyAll = nil

	statement, err := models.Db.Prepare(
		`SELECT k.id AS "key_id",
            k.name,
            k.description,
            k.key_type_id,
            kt.name AS "key_type_name",
            s.id AS "service_id",
            s.name AS "service_name"
        FROM keys k
        LEFT OUTER JOIN key_types kt ON k.key_type_id = kt.id
        LEFT OUTER JOIN service_keys sk ON k.id = sk.key_id
        LEFT OUTER JOIN services s ON sk.service_id = s.id
        ORDER BY s.id`)
	if err != nil {
		return nil, err
	}

	rows, err := statement.Query()
	if err != nil {
		return nil, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	for rows.Next() {
		err := insertService(rows)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err = rows.Err(); err != nil {
		return keyAll, err
	}

	return keyAll, nil
}

func insertService(rows *sql.Rows) error {
	var id int
	var name string
	var description string
	var typeId int
	var typeName sql.NullString
	var serviceId sql.NullInt64
	var serviceName sql.NullString

	err := rows.Scan(&id, &name, &description, &typeId, &typeName, &serviceId, &serviceName)
	if err != nil {
		return err
	}

	if oldKeyId != id {
		var kAll KeyAll

		kAll.Id = id
		kAll.Name = name
		kAll.Description = description
		kAll.KeyTypeId = typeId
		kAll.KeyTypeName = typeName.String

		keyAll = append(keyAll, kAll)

		oldKeyId = id
	}

	if serviceId.Valid {
		var newService service

		newService.Id = int(serviceId.Int64)
		newService.Name = serviceName.String

		newServ := &keyAll[len(keyAll)-1]
		newServ.Services = append(newServ.Services, newService)
	}

	return nil
}
