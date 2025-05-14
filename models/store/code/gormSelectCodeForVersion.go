package code

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
	"strings"
)

type GormCode struct {
	Id    int    `json:"id"`
	UUID  string `json:"uuid"`
	Type  string `json:"type"`
	Order int    `json:"descLong"`
	Value string `json:"value"`
}

func GormSelectCodeForVersion(scriptUUID string, versionId int) ([]GormCode, error) {
	var query strings.Builder

	query.WriteString("SELECT c.id, ")
	query.WriteString("       BIN_TO_UUID(c.uuid) AS 'uuid', ")
	query.WriteString("       c.type, ")
	query.WriteString("       vc.order, ")
	query.WriteString("       c.value ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("INNER JOIN versions AS v ")
	query.WriteString("ON s.id = v.script_id ")
	query.WriteString("INNER JOIN version_codes AS vc ")
	query.WriteString("ON v.id = vc.version_id ")
	query.WriteString("INNER JOIN codes AS c ")
	query.WriteString("ON vc.code_id = c.id ")
	query.WriteString("WHERE s.UUID = UUID_TO_BIN(?) ")
	query.WriteString("AND v.id = ? ")
	query.WriteString("ORDER BY vc.order ")

	fmt.Println("GormSelectCodeForVersion - query", query.String())

	var output []GormCode

	errGorm := store.GormDB.Raw(query.String(), scriptUUID, versionId).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GormSelectCodeForVersion - ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	fmt.Println("GormSelectCodeForVersion - output", output)

	return output, nil
}
