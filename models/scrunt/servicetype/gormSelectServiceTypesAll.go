package servicetype

import (
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/scrunt"
	"strings"
)

type GormServiceType struct {
	Uuid      string
	Directory string
	Filename  string
	Name      string
	DescShort string
	DescLong  string
	Local     bool
}

func GormSelectServiceTypesAll() ([]GormServiceType, error) {
	var query strings.Builder

	query.WriteString("SELECT st.uuid, ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename, ")
	query.WriteString("       st.name, ")
	query.WriteString("       st.desc_short, ")
	query.WriteString("       st.desc_long, ")
	query.WriteString("       st.local ")
	query.WriteString("FROM service_types AS st ")
	query.WriteString("LEFT OUTER JOIN icons AS i ")
	query.WriteString("ON st.icon_code = i.code ")

	var output []GormServiceType
	result := scrunt.GormDB.Raw(query.String()).Scan(&output)

	if result.Error != nil {
		return nil, result.Error
	}

	return output, nil
}
