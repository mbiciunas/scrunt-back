package script

import (
	"fmt"
	"scrunt-back/models/scrunt"
	"strings"
)

type GormScript struct {
	Id        int    `json:"id"`
	Directory string `json:"directory"`
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	DescShort string `json:"descShort"`
	DescLong  string `json:"descLong"`
	Created   string `json:"created"`
	Download  int    `json:"download"`
}

func GormSelectScript(id int) (GormScript, error) {
	var query strings.Builder

	query.WriteString("SELECT s.id, ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename, ")
	query.WriteString("       s.name, ")
	query.WriteString("       s.desc_short, ")
	query.WriteString("       s.desc_long, ")
	query.WriteString("       s.created, ")
	query.WriteString("       s.download ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("LEFT OUTER JOIN icons AS i ")
	query.WriteString("ON s.icon_id = i.id ")
	query.WriteString("WHERE s.id = ? ")

	//fmt.Println("store.script.gormSelectScript - query: ", query.String())

	var output GormScript

	errGorm := scrunt.GormDB.Raw(query.String(), id).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("store.script.gormSelectScript - errGorm: ", errGorm)
		return output, errGorm.Error
	}

	//fmt.Println("store.script.gormSelectScript - output: ", output)

	return output, nil
}
