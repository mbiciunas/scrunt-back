package version

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/scrunt"
	"strings"
)

type GormVersionAll struct {
	Id       int    `json:"id"`
	ScriptId int    `json:"scriptId"`
	Created  string `json:"created"`
	Major    int    `json:"major"`
	Minor    int    `json:"minor"`
	Patch    int    `json:"patch"`
	Save     int    `json:"save"`
	UUID     string `json:"uuid"`
	Change   string `json:"change"`
}

func GormSelectVersionsAll(scriptUUID string) ([]GormVersionAll, error) {
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
	query.WriteString("ORDER BY v.created DESC ")

	//fmt.Println(">>>", query.String(), "<<<")

	var output []GormVersionAll

	errGorm := scrunt.GormDB.Raw(query.String(), scriptUUID).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GORM ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	//fmt.Println("GORM selectVersionsAll", output)

	return output, nil
}
