package script

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
	"strings"
)

type GormScriptAll struct {
	Id        int     `json:"id"`
	Uuid      string  `json:"uuid"`
	Directory string  `json:"directory"`
	Filename  string  `json:"filename"`
	Name      string  `json:"name"`
	DescShort string  `json:"desc_short"`
	DescLong  string  `json:"desc_long"`
	Created   string  `json:"created"`
	Rating    float32 `json:"rating"`
	Tag       string  `json:"tag"`
}

func GormSelectScriptsAll() ([]GormScriptAll, error) {
	var query strings.Builder

	query.WriteString("SELECT s.id, ")
	query.WriteString("       BIN_TO_UUID(s.uuid) AS uuid, ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename, ")
	query.WriteString("       s.name, ")
	query.WriteString("       s.desc_short, ")
	query.WriteString("       s.desc_long, ")
	query.WriteString("       s.created, ")
	query.WriteString("       AVG(r.amount) AS rating, ")
	query.WriteString("       (SELECT GROUP_CONCAT(t2.name) ")
	query.WriteString("        FROM tags AS t2")
	query.WriteString("        LEFT OUTER JOIN script_tags AS st2 ")
	query.WriteString("        ON st2.tag_id = t2.id ")
	query.WriteString("        WHERE s.id = st2.script_id) AS tag ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("LEFT OUTER JOIN icons AS i ")
	query.WriteString("ON s.icon_code = i.code ")
	query.WriteString("LEFT OUTER JOIN ratings AS r ")
	query.WriteString("ON s.id = r.script_id ")
	query.WriteString("LEFT OUTER JOIN script_tags AS st ")
	query.WriteString("ON s.id = st.script_id ")
	query.WriteString("LEFT OUTER JOIN tags AS t ")
	query.WriteString("ON st.tag_id = t.id ")
	query.WriteString("WHERE t.name LIKE \"%\" ")
	query.WriteString("GROUP BY s.id ")

	var output []GormScriptAll
	errGorm := store.GormDB.Raw(query.String()).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GORM ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	return output, nil
}
