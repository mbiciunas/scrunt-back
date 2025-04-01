package icon

import (
	_ "github.com/mattn/go-sqlite3"
	"scrunt-back/models"
	"scrunt-back/models/scrunt"
	"strings"
)

//type GormIconAll struct {
//	Id        int
//	Code      string
//	Directory string
//	Filename  string
//}

func GormSelectIconsAll() ([]models.Icon, error) {
	var query strings.Builder

	query.WriteString("SELECT i.id, ")
	query.WriteString("       i.code, ")
	query.WriteString("       i.directory, ")
	query.WriteString("       i.filename ")
	query.WriteString("FROM icons AS i ")

	var output []models.Icon
	errGorm := scrunt.GormDB.Raw(query.String()).Scan(&output)

	if errGorm.Error != nil {
		return nil, errGorm.Error
	}

	return output, nil
}
