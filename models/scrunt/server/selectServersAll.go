package server

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models"
)

type credential struct {
	Id       int
	Name     string
	CredType string
	ScId     int
}

type ServerAll struct {
	Id          int
	Name        string
	Address     string
	Credentials []credential
}

var serverAll []ServerAll
var oldServerId int

func SelectServersAll() ([]ServerAll, error) {
	serverAll = nil

	statement, err := models.Db.Prepare(
		`SELECT s.id AS "serv_id",
			s.name,
			s.address,
			c.id AS "cred_id",
			c.name,
			c.type,
            sc.id AS "sc_id"
		FROM servers s
		LEFT JOIN server_credentials sc ON s.id = sc.server_id
		LEFT JOIN credentials c  ON sc.credential_id = c.id
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
		err := insertServer(rows)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err = rows.Err(); err != nil {
		return serverAll, err
	}

	return serverAll, nil
}

func insertServer(rows *sql.Rows) error {
	var serverId int
	var serverName string
	var serverAddress string
	var credentialId sql.NullInt64
	var credentialName sql.NullString
	var credentialType sql.NullString
	var scId sql.NullInt64

	err := rows.Scan(&serverId, &serverName, &serverAddress, &credentialId, &credentialName, &credentialType, &scId)
	if err != nil {
		return err
	}

	if oldServerId != serverId {
		var servAll ServerAll

		servAll.Id = serverId
		servAll.Name = serverName
		servAll.Address = serverAddress

		serverAll = append(serverAll, servAll)

		oldServerId = serverId
	}

	if credentialId.Valid {
		var newCredential credential

		newCredential.Id = int(credentialId.Int64)
		newCredential.Name = credentialName.String
		newCredential.CredType = credentialType.String
		newCredential.ScId = int(scId.Int64)

		newServ := &serverAll[len(serverAll)-1]
		newServ.Credentials = append(newServ.Credentials, newCredential)
	}

	return nil
}
