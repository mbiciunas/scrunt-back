package version

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
	"strings"
)

func GormSelectVersionNewest(id int) (GormVersionAll, error) {
	var query strings.Builder

	query.WriteString("SELECT id, ")
	query.WriteString("       script_id, ")
	query.WriteString("       created, ")
	query.WriteString("       major, ")
	query.WriteString("       minor, ")
	query.WriteString("       patch, ")
	query.WriteString("       BIN_TO_UUID(uuid) AS uuid, ")
	query.WriteString("       `change` ")
	query.WriteString("FROM versions ")
	query.WriteString("WHERE script_id = ? ")
	query.WriteString("AND created = (SELECT max(created) from versions WHERE script_id = ?) ")

	fmt.Println(">>>", query.String(), "<<<")

	var output GormVersionAll

	errGorm := store.GormDB.Raw(query.String(), id, id).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GORM ERROR Raw: ", errGorm)
		return output, errGorm.Error
	}

	//fmt.Println("GORM selectVersionsAll", output)

	return output, nil
}
