package credential

import (
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type Credential struct {
	Id         int
	Name       string
	Type       string
	Credential string
}

func SelectCredential(id int) (Credential, error) {
	statement, err := scrunt.Db.Prepare(`SELECT id, name, type, credential FROM credentials WHERE id = $1`)
	if err != nil {
		return Credential{}, err
	}

	rows, err := statement.Query(id)
	if err != nil {
		return Credential{}, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	var credential Credential

	for rows.Next() {
		//var cred credential

		err := rows.Scan(&credential.Id, &credential.Name, &credential.Type, &credential.Credential)
		if err != nil {
			return credential, err
		}
	}
	if err = rows.Err(); err != nil {
		return credential, err
	}

	return credential, nil
}
