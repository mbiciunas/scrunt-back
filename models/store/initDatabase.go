package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDB() error {
	database, err := sql.Open("mysql", "root:M@rk8478@tcp(127.0.0.1:3306)/store")
	if err != nil {
		return err
	}

	Db = database

	return Db.Ping()
}
