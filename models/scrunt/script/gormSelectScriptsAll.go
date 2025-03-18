package script

import (
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/scrunt"
	"strings"
)

type GormScriptAll struct {
	Id        int
	Directory string
	Filename  string
	Name      string
	DescShort string
	DescLong  string
	Created   string
	Tag       string
}

func GormSelectScriptsAll() ([]GormScriptAll, error) {
	var query strings.Builder

	query.WriteString("SELECT s.id, ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename, ")
	query.WriteString("       s.name, ")
	query.WriteString("       s.desc_short, ")
	query.WriteString("       s.desc_long, ")
	query.WriteString("       s.created, ")
	query.WriteString("       (SELECT GROUP_CONCAT(t2.name) ")
	query.WriteString("        FROM tags AS t2")
	query.WriteString("        LEFT OUTER JOIN script_tags AS st2 ")
	query.WriteString("        ON st2.tag_id = t2.id ")
	query.WriteString("        WHERE s.id = st2.script_id) AS tag ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("LEFT OUTER JOIN icons AS i ")
	query.WriteString("ON s.icon_id = i.id ")
	query.WriteString("GROUP BY s.id ")

	var output []GormScriptAll
	errGorm := scrunt.GormDB.Raw(query.String()).Scan(&output)

	if errGorm.Error != nil {
		return nil, errGorm.Error
	}

	return output, nil
}
