package service

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type Service struct {
	Id            int
	Name          string
	Description   string
	Address       string
	Port          int
	ServiceTypeId int
	CloudId       int
}

func SelectService(id int) (Service, error) {
	statement, err := scrunt.Db.Prepare(`SELECT id, name, description, address, port, service_type_id, cloud_id FROM services WHERE id = $1`)
	if err != nil {
		fmt.Println("SelectService - prepare", err)
		return Service{}, err
	}

	rows, err := statement.Query(id)
	if err != nil {
		fmt.Println("SelectService - query", err)
		return Service{}, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	var service Service

	for rows.Next() {
		err := rows.Scan(&service.Id, &service.Name, &service.Description, &service.Address, &service.Port, &service.ServiceTypeId, &service.CloudId)
		if err != nil {
			return service, err
		}
	}
	if err = rows.Err(); err != nil {
		return service, err
	}

	return service, nil
}
