package runtime

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Output struct {
	Id          int
	OutputType  int
	OutputValue string
}

var output []Output

func SelectOutputByRunId(runId int, id int) ([]Output, error) {
	output = nil
	statement, err := db.Prepare(`SELECT id, type, value FROM outputs WHERE run_id = $1 and id > $2`)
	if err != nil {
		fmt.Println("SelectOutputByRunId - prepare", err)
		return output, err
	}

	rows, err := statement.Query(runId, id)
	if err != nil {
		fmt.Println("SelectOutputByRunId - query", err)
		return output, err
	}
	defer func() {
		ferr := rows.Close()
		if ferr != nil {
			err = ferr
		}
	}()

	for rows.Next() {
		err := insertOutput(rows)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	fmt.Println("Final", output)

	return output, nil
}

func insertOutput(rows *sql.Rows) error {
	var id int
	var outputType int
	var outputValue string

	fmt.Println("models.selectOutputsByRunId", "Insert Output")
	err := rows.Scan(&id, &outputType, &outputValue)
	if err != nil {
		fmt.Println("models.selectOutputsByRunId", "Error", err)
		return err
	}

	var outputRow Output

	outputRow.Id = id
	outputRow.OutputType = outputType
	outputRow.OutputValue = outputValue

	fmt.Println(outputRow)

	output = append(output, outputRow)

	return nil
}
