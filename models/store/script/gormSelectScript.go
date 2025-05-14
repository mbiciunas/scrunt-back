package script

import (
	"fmt"
	"scrunt-back/models/store"
	"strings"
)

type GormScript struct {
	Id        int    `json:"id"`
	Uuid      string `json:"uuid"`
	IconCode  string `json:"iconCode"`
	Directory string `json:"directory"`
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	DescShort string `json:"descShort"`
	DescLong  string `json:"descLong"`
	Created   string `json:"created"`
	Download  int    `json:"download"`
}

func GormSelectScript(scriptUUID string) (GormScript, error) {
	var query strings.Builder

	query.WriteString("SELECT s.id, ")
	query.WriteString("       BIN_TO_UUID(s.uuid) AS uuid, ")
	query.WriteString("       i.code AS 'icon_code', ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename, ")
	query.WriteString("       s.name, ")
	query.WriteString("       s.desc_short, ")
	query.WriteString("       s.desc_long, ")
	query.WriteString("       s.created, ")
	query.WriteString("       s.download ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("LEFT OUTER JOIN icons AS i ")
	query.WriteString("ON s.icon_code = i.code ")
	query.WriteString("WHERE s.uuid = UUID_TO_BIN(?) ")

	var output GormScript

	errGorm := store.GormDB.Raw(query.String(), scriptUUID).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("store.script.gormSelectScript - errGorm: ", errGorm)
		return output, errGorm.Error
	}

	return output, nil
}
