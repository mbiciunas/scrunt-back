package cloud

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

type CloudAll struct {
	Id          int
	Name        string
	Description string
	CloudTypeId int
}

var cloudAll []CloudAll

func SelectCloudAll() ([]CloudAll, error) {
	cloudAll = nil

	fmt.Println("models.selectCloudsAll - SelectCloudAll")

	statement, err := scrunt.Db.Prepare(
		`SELECT c.id,
            c.name,
            c.description,
            c.cloud_type_id
        FROM clouds c
        ORDER BY c.name`)
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
		err := insertCloud(rows)
		if err != nil {
			fmt.Println(err)
		}
	}
	if err = rows.Err(); err != nil {
		return cloudAll, err
	}

	return cloudAll, nil
}

func insertCloud(rows *sql.Rows) error {

	var id int
	var name string
	var description string
	var cloudTypeId sql.NullInt64
	var cloud CloudAll

	fmt.Println("models.selectServiceTypesAll - insertServiceType")

	err := rows.Scan(&id, &name, &description, &cloudTypeId)
	if err != nil {
		return err
	}

	cloud.Id = id
	cloud.Name = name
	cloud.Description = description
	cloud.CloudTypeId = int(cloudTypeId.Int64)

	cloudAll = append(cloudAll, cloud)

	return nil
}
