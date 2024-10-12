package server

import (
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type Server struct {
	Id      int
	Name    string
	Address string
}

func SelectServer(id int) (Server, error) {
	statement, err := scrunt.Db.Prepare(`SELECT id, name, address FROM servers WHERE id = $1`)
	if err != nil {
		return Server{}, err
	}

	rows, err := statement.Query(id)
	if err != nil {
		return Server{}, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	var server Server

	for rows.Next() {
		//var cred credential

		err := rows.Scan(&server.Id, &server.Name, &server.Address)
		if err != nil {
			return server, err
		}
	}
	if err = rows.Err(); err != nil {
		return server, err
	}

	return server, nil
}
