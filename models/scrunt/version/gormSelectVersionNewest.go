package version

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/scrunt"
	"strings"
)

func GormSelectVersionNewest(scriptUUID string) (GormVersionAll, error) {
	var query strings.Builder

	query.WriteString("SELECT v.id, ")
	query.WriteString("       v.script_id, ")
	query.WriteString("       v.created, ")
	query.WriteString("       v.major, ")
	query.WriteString("       v.minor, ")
	query.WriteString("       v.patch, ")
	query.WriteString("       v.save, ")
	query.WriteString("       v.uuid, ")
	query.WriteString("       v.change ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("INNER JOIN versions AS v ")
	query.WriteString("ON s.id = v.script_id ")
	query.WriteString("WHERE s.uuid = ? ")
	query.WriteString("AND v.created = (SELECT max(v1.created) from versions AS v1 WHERE v1.script_id = s.id) ")

	var output GormVersionAll

	result := scrunt.GormDB.Raw(query.String(), scriptUUID).Scan(&output)

	if result.Error != nil {
		fmt.Println("GORM ERROR Raw: ", result.Error)
		return output, result.Error
	}

	return output, nil
}
