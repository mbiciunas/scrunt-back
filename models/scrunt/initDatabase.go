package scrunt

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *sql.DB
var GormDB *gorm.DB

func InitDB() error {
	database, err := sql.Open("sqlite3", ".scrunt/database/scrunt.db")
	if err != nil {
		return err
	}

	Db = database

	return Db.Ping()
}

func InitGorm() error {
	database, err := gorm.Open(sqlite.Open(".scrunt/database/gorm.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	GormDB = database

	return err

}
