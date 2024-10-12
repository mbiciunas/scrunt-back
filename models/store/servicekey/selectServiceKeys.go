package servicekey

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/store"
)

type ServiceKey struct {
	ServiceKeyId int
	KeyId        int
	Name         string
	Description  string
	KeyTypeId    int
	KeyTypeName  string
}

var serviceKeys []ServiceKey

func SelectServiceKeys(serviceId int) ([]ServiceKey, error) {
	serviceKeys = nil

	fmt.Println("models.selectServiceKeys")
	statement, err := store.Db.Prepare(
		`SELECT sk.id AS "service_key_id",
			k.id AS "key_id",
			k.name,
			k.description,
			k.key_type_id,
			kt.name AS "key_type_name"
		FROM service_keys sk
		INNER JOIN keys k ON k.id = sk.key_id
		INNER JOIN key_types kt ON kt.id = k.key_type_id
		WHERE sk.service_id = $1`)
	if err != nil {
		fmt.Println("models.selectScriptServices", "Error", err)
		return nil, err
	}

	fmt.Println("models.selectScriptServices", "Execute query")
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
		err := insertServiceKeys(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return serviceKeys, err
	}

	fmt.Println("Final", serviceKeys)

	return serviceKeys, nil
}

func insertServiceKeys(rows *sql.Rows) error {
	var serviceKeyId int
	var keyId int
	var name string
	var description string
	var keyTypeId int
	var keyTypeName string

	fmt.Println("models.selectServiceKeys.insertServiceKeys", "Insert ServiceKeys")
	err := rows.Scan(&serviceKeyId, &keyId, &name, &description, &keyTypeId, &keyTypeName)
	if err != nil {
		fmt.Println("models.selectServiceKeys.insertServiceKeys", "Error", err)
		return err
	}

	var serviceKey ServiceKey

	serviceKey.ServiceKeyId = serviceKeyId
	serviceKey.KeyId = keyId
	serviceKey.Name = name
	serviceKey.Description = description
	serviceKey.KeyTypeId = keyTypeId
	serviceKey.KeyTypeName = keyTypeName

	fmt.Println(serviceKey)

	serviceKeys = append(serviceKeys, serviceKey)

	fmt.Println(serviceKeys)

	return nil
}
