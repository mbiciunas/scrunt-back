package issue

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"scrunt-back/models/store"
	"strings"
)

type GormIssue struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Alias       string `json:"alias"`
	Comments    int    `json:"comments"`
}

func GormSelectIssues(scriptId int) ([]GormIssue, error) {
	var query strings.Builder

	query.WriteString("SELECT i.id, ")
	query.WriteString("       i.date, ")
	query.WriteString("       i.title, ")
	query.WriteString("       i.description, ")
	query.WriteString("       i.type, ")
	query.WriteString("       i.status, ")
	query.WriteString("       u.alias, ")
	query.WriteString("	     count(ic.id) AS 'comments' ")
	query.WriteString("FROM issues AS i ")
	query.WriteString("INNER JOIN users AS u ")
	query.WriteString("ON i.user_id = u.id ")
	query.WriteString("LEFT OUTER JOIN issue_comments AS ic ")
	query.WriteString("ON i.id = ic.issue_id ")
	query.WriteString("WHERE i.script_id = ? ")
	query.WriteString("GROUP BY i.id ")

	fmt.Println("GormSelectIssues - query", query.String())

	var output []GormIssue

	errGorm := store.GormDB.Raw(query.String(), scriptId).Scan(&output)

	if errGorm.Error != nil {
		fmt.Println("GormSelectIssues - ERROR Raw: ", errGorm)
		return nil, errGorm.Error
	}

	fmt.Println("GormSelectIssues - output", output)

	return output, nil
}
