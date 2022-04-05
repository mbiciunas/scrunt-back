package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type server struct {
	Id      int
	Name    string
	Address string
	ScId    int
}

type CredentialAll struct {
	Id       int
	Name     string
	CredType string
	Servers  []server
}

var credentialAll []CredentialAll
var oldCredentialId int

func SelectCredentialsAll() ([]CredentialAll, error) {
	credentialAll = nil

	statement, err := db.Prepare(
		`SELECT c.id AS "cred_id",
			c.name,
			c.type,
			s.id AS "serv_id",
			s.name,
			s.address,
            sc.id AS "sc_id"
		FROM credentials c 
		LEFT JOIN server_credentials sc ON c.id = sc.credential_id 
		LEFT JOIN servers s  ON sc.server_id  = s.id
		ORDER BY c.id`)
	if err != nil {
		return nil, err
	}

	rows, err := statement.Query()
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
		err := insertCredential(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return credentialAll, err
	}

	fmt.Println("Final", credentialAll)

	return credentialAll, nil
}

func insertCredential(rows *sql.Rows) error {
	var credentialId int
	var credentialName string
	var credentialType string
	var serverId sql.NullInt64
	var serverName sql.NullString
	var serverAddress sql.NullString
	var scId sql.NullInt64

	err := rows.Scan(&credentialId, &credentialName, &credentialType, &serverId, &serverName, &serverAddress, &scId)
	if err != nil {
		return err
	}

	if oldCredentialId != credentialId {
		var credAll CredentialAll

		credAll.Id = credentialId
		credAll.Name = credentialName
		credAll.CredType = credentialType

		fmt.Println(credAll)

		credentialAll = append(credentialAll, credAll)

		fmt.Println(credentialAll)

		oldCredentialId = credentialId
	}

	if serverId.Valid {
		var newServer server

		newServer.Id = int(serverId.Int64)
		newServer.Name = serverName.String
		newServer.Address = serverAddress.String
		newServer.ScId = int(serverId.Int64)

		fmt.Println("newServer", newServer)

		newCred := &credentialAll[len(credentialAll)-1]
		newCred.Servers = append(newCred.Servers, newServer)
	}

	return nil
}
