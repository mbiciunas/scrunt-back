package version

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/scrunt"
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
	query.WriteString("       save, ")
	query.WriteString("       uuid, ")
	query.WriteString("       change ")
	query.WriteString("FROM versions ")
	query.WriteString("WHERE script_id = ? ")
	query.WriteString("AND created = (SELECT max(created) from versions WHERE script_id = ?) ")

	var output GormVersionAll

	result := scrunt.GormDB.Raw(query.String(), id, id).Scan(&output)

	if result.Error != nil {
		fmt.Println("GORM ERROR Raw: ", result.Error)
		return output, result.Error
	}

	return output, nil
}
