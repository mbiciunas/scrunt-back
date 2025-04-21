package version

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
	"strings"
)

type GormVersionAll struct {
	Id       int    `json:"id"`
	ScriptId int    `json:"scriptId"`
	Created  string `json:"created"`
	Major    int    `json:"major"`
	Minor    int    `json:"minor"`
	Patch    int    `json:"patch"`
	UUID     string `json:"uuid"`
	Changes  string `json:"changes"`
}

func GormSelectVersionsAll(scriptId int) ([]GormVersionAll, error) {
	var query strings.Builder

	query.WriteString("SELECT v.id, ")
	query.WriteString("       v.script_id, ")
	query.WriteString("       v.created, ")
	query.WriteString("       v.major, ")
	query.WriteString("       v.minor, ")
	query.WriteString("       v.patch, ")
	query.WriteString("       BIN_TO_UUID(v.uuid) AS uuid, ")
	query.WriteString("       v.changes ")
	query.WriteString("FROM versions AS v ")
	query.WriteString("WHERE v.script_id = ? ")
	query.WriteString("ORDER BY v.created DESC ")

	var output []GormVersionAll

	errGorm := store.GormDB.Raw(query.String(), scriptId).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GORM ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	//fmt.Println("GORM selectVersionsAll", output)

	return output, nil
}
