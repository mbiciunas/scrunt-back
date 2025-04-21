package script

import (
	"fmt"
	"scrunt-back/models/scrunt"
	"strings"
)

type GormScript struct {
	Id        int    `json:"id"`
	Uuid      string `json:"uuid"`
	Directory string `json:"directory"`
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	DescShort string `json:"descShort"`
	DescLong  string `json:"descLong"`
	Source    uint   `json:"source"`
	Parent    string `json:"parent"`
	Created   string `json:"created"`
}

func GormSelectScript(id int) (GormScript, error) {
	var query strings.Builder

	query.WriteString("SELECT s.id, ")
	query.WriteString("       s.uuid, ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename, ")
	query.WriteString("       s.name, ")
	query.WriteString("       s.desc_short, ")
	query.WriteString("       s.desc_long, ")
	query.WriteString("       s.source, ")
	query.WriteString("       s.parent, ")
	query.WriteString("       s.created ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("LEFT OUTER JOIN icons AS i ")
	query.WriteString("ON s.icon_code = i.code ")
	query.WriteString("WHERE s.id = ? ")

	var output GormScript

	errGorm := scrunt.GormDB.Raw(query.String(), id).Scan(&output)

	fmt.Println("models.scrunt.script.gormSelectScript - output", output)

	if errGorm.Error != nil {
		return output, errGorm.Error
	}

	return output, nil
}
