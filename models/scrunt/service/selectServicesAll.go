package service

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

type project struct {
	Id   int
	Name string
}

type ServiceAll struct {
	Id              int
	Name            string
	Description     string
	Address         string
	Port            int
	ServiceTypeId   int
	ServiceTypeName string
	CloudId         int
	CloudName       string
	Projects        []project
}

var serviceAll []ServiceAll
var oldServiceId int

func SelectServicesAll() ([]ServiceAll, error) {
	serviceAll = nil

	statement, err := scrunt.Db.Prepare(
		`SELECT s.id AS "serv_id",
            s.name,
            s.description,
            s.address,
            s.port,
            s.service_type_id,
            st.name AS "service_type_name",
            s.cloud_id,
            c.name AS "cloud_name",
            p.id AS "project_id",
            p.name AS "project_name"
        FROM services s
        LEFT OUTER JOIN service_types st ON s.service_type_id = st.id
        LEFT OUTER JOIN clouds c ON s.cloud_id = c.id
        LEFT OUTER JOIN service_projects sp ON s.id = sp.service_id
        LEFT OUTER JOIN projects p  ON sp.project_id = p.id
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
		return serviceAll, err
	}

	return serviceAll, nil
}

func insertService(rows *sql.Rows) error {

	var id int
	var name string
	var description string
	var address string
	var port int
	var typeId int
	var typeName sql.NullString
	var cloudId sql.NullInt64
	var cloudName sql.NullString
	var projectId sql.NullInt64
	var projectName sql.NullString

	err := rows.Scan(&id, &name, &description, &address, &port, &typeId, &typeName, &cloudId, &cloudName, &projectId, &projectName)
	if err != nil {
		return err
	}

	if oldServiceId != id {
		var servAll ServiceAll

		servAll.Id = id
		servAll.Name = name
		servAll.Description = description
		servAll.Address = address
		servAll.Port = port
		servAll.ServiceTypeId = typeId
		servAll.ServiceTypeName = typeName.String
		servAll.CloudId = int(cloudId.Int64)
		servAll.CloudName = cloudName.String

		serviceAll = append(serviceAll, servAll)

		oldServiceId = id
	}

	if projectId.Valid {
		var newProject project

		newProject.Id = int(projectId.Int64)
		newProject.Name = projectName.String

		newServ := &serviceAll[len(serviceAll)-1]
		newServ.Projects = append(newServ.Projects, newProject)
	}

	return nil
}
