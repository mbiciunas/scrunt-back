package keytype

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models/scrunt"
)

type KeyTypeAll struct {
	Id   int
	Name string
	Icon string
}

var keyTypeAll []KeyTypeAll

func SelectKeyTypesAll() ([]KeyTypeAll, error) {
	keyTypeAll = nil

	fmt.Println("models.selectKeyTypesAll - SelectKeyTypesAll")

	statement, err := scrunt.Db.Prepare(
		`SELECT k.id,
            k.name,
            k.icon
        FROM key_types k
        ORDER BY k.name`)
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
		err := insertKeyType(rows)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err = rows.Err(); err != nil {
		return keyTypeAll, err
	}

	return keyTypeAll, nil
}

func insertKeyType(rows *sql.Rows) error {

	var id int
	var name string
	var icon sql.NullString
	var kTypeAll KeyTypeAll

	fmt.Println("models.selectServiceTypesAll - insertServiceType")

	err := rows.Scan(&id, &name, &icon)
	if err != nil {
		return err
	}

	kTypeAll.Id = id
	kTypeAll.Name = name
	kTypeAll.Icon = icon.String

	keyTypeAll = append(keyTypeAll, kTypeAll)

	return nil
}
