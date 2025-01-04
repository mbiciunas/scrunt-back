package rating

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
	"strings"
)

type GormRating struct {
	Type   string  `json:"type"`
	Rating float32 `json:"rating"`
}

func GormSelectRatingSummary(id int) ([]GormRating, error) {
	var query strings.Builder

	query.WriteString("SELECT type, ")
	query.WriteString("       SUM(amount) / COUNT(type) AS rating ")
	query.WriteString("FROM ratings ")
	query.WriteString("WHERE script_id = ? ")
	query.WriteString("GROUP BY type ")

	fmt.Println(">>>", query.String(), "<<<")

	var output []GormRating

	errGorm := store.GormDB.Raw(query.String(), id).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GORM ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	fmt.Println("GormSelectRatingSummary", output)

	return output, nil
}
