package script

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
)

type ScriptAll struct {
	Id        int
	Name      string
	DescShort string
	DescLong  string
}

var scriptAll []ScriptAll

func SelectScriptsAll() ([]ScriptAll, error) {
	scriptAll = nil

	fmt.Println("models.selectScriptsAll")
	statement, err := store.Db.Prepare(
		`SELECT s.id,
			s.name,
			s.description_short,
			s.description_long
		FROM scripts s`)
	if err != nil {
		fmt.Println("models.selectScriptsAll", "Error", err)
		return nil, err
	}

	fmt.Println("models.selectScriptsAll", "Execute query")
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
		err := insertScript(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return scriptAll, err
	}

	fmt.Println("Final", scriptAll)

	return scriptAll, nil
}

func insertScript(rows *sql.Rows) error {
	var id int
	var name string
	var descriptionShort string
	var descriptionLong string

	fmt.Println("models.selectScriptsAll", "Insert Script")
	err := rows.Scan(&id, &name, &descriptionShort, &descriptionLong)
	if err != nil {
		fmt.Println("models.selectScriptsAll", "Error", err)
		return err
	}

	var scriptRow ScriptAll

	scriptRow.Id = id
	scriptRow.Name = name
	scriptRow.DescShort = descriptionShort
	scriptRow.DescLong = descriptionLong

	fmt.Println(scriptRow)

	scriptAll = append(scriptAll, scriptRow)

	fmt.Println(scriptAll)

	return nil
}
