package servicetype

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type ServiceTypeAll struct {
	Id   int
	Name string
	Icon string
}

var serviceTypeAll []ServiceTypeAll

func SelectServiceTypesAll() ([]ServiceTypeAll, error) {
	serviceTypeAll = nil

	fmt.Println("models.selectServiceTypesAll - SelectServiceTypesAll")

	statement, err := scrunt.Db.Prepare(
		`SELECT s.id,
            s.name,
            s.icon
        FROM service_types s
        ORDER BY s.name`)
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
		err := insertServiceType(rows)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err = rows.Err(); err != nil {
		return serviceTypeAll, err
	}

	return serviceTypeAll, nil
}

func insertServiceType(rows *sql.Rows) error {

	var id int
	var name string
	var icon sql.NullString
	var servTypeAll ServiceTypeAll

	fmt.Println("models.selectServiceTypesAll - insertServiceType")

	err := rows.Scan(&id, &name, &icon)
	if err != nil {
		return err
	}

	servTypeAll.Id = id
	servTypeAll.Name = name
	servTypeAll.Icon = icon.String

	serviceTypeAll = append(serviceTypeAll, servTypeAll)

	return nil
}
