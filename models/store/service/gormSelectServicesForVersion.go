package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
	"strings"
)

type GormService struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	DescShort string `json:"descShort"`
	DescLong  string `json:"descLong"`
	Directory string `json:"directory"`
	Filename  string `json:"filename"`
}

func GormSelectServicesForVersion(id int) ([]GormService, error) {
	var query strings.Builder

	query.WriteString("SELECT s.id, ")
	query.WriteString("       s.name, ")
	query.WriteString("       s.desc_short, ")
	query.WriteString("       s.desc_long, ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename ")
	query.WriteString("FROM services AS s ")
	query.WriteString("INNER JOIN icons AS i ")
	query.WriteString("ON s.icon_id = i.id ")
	query.WriteString("INNER JOIN version_services AS vs ")
	query.WriteString("ON s.id = vs.service_id ")
	query.WriteString("WHERE vs.version_id = ? ")

	fmt.Println(">>>", query.String(), "<<<")

	var output []GormService

	errGorm := store.GormDB.Raw(query.String(), id).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GORM ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	fmt.Println("GormSelectServicesForVersion", output)

	return output, nil
}
