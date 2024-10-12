package project

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type ProjectAll struct {
	Id          int
	Name        string
	Description string
	Icon        string
}

var projectAll []ProjectAll

func SelectProjectAll() ([]ProjectAll, error) {
	projectAll = nil

	fmt.Println("models.selectProjectsAll - SelectProjectAll")

	statement, err := scrunt.Db.Prepare(
		`SELECT p.id,
            p.name,
            p.description,
            p.icon
        FROM projects p
        ORDER BY p.name`)
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
		err := insertProject(rows)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err = rows.Err(); err != nil {
		return projectAll, err
	}

	return projectAll, nil
}

func insertProject(rows *sql.Rows) error {

	var id int
	var name string
	var description string
	var icon sql.NullString
	var project ProjectAll

	fmt.Println("models.selectServiceTypesAll - insertServiceType")

	err := rows.Scan(&id, &name, &description, &icon)
	if err != nil {
		return err
	}

	project.Id = id
	project.Name = name
	project.Description = description
	project.Icon = icon.String

	projectAll = append(projectAll, project)

	return nil
}
