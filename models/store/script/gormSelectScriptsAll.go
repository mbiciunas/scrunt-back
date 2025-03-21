package script

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
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
	Rating    float32
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
	query.WriteString("       AVG(r.amount) AS rating, ")
	query.WriteString("       (SELECT GROUP_CONCAT(t2.name) ")
	query.WriteString("        FROM tags AS t2")
	query.WriteString("        LEFT OUTER JOIN script_tags AS st2 ")
	query.WriteString("        ON st2.tag_id = t2.id ")
	query.WriteString("        WHERE s.id = st2.script_id) AS tag ")
	query.WriteString("FROM scripts AS s ")
	query.WriteString("LEFT OUTER JOIN icons AS i ")
	query.WriteString("ON s.icon_id = i.id ")
	query.WriteString("LEFT OUTER JOIN ratings AS r ")
	query.WriteString("ON s.id = r.script_id ")
	query.WriteString("LEFT OUTER JOIN script_tags AS st ")
	query.WriteString("ON s.id = st.script_id ")
	query.WriteString("LEFT OUTER JOIN tags AS t ")
	query.WriteString("ON st.tag_id = t.id ")
	query.WriteString("WHERE t.name LIKE \"%\" ")
	query.WriteString("GROUP BY s.id ")

	//fmt.Println(">>>", query.String(), "<<<")

	var output []GormScriptAll
	errGorm := store.GormDB.Raw(query.String()).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GORM ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	//fmt.Println("GORM selectScriptsAll", output)

	return output, nil
}
