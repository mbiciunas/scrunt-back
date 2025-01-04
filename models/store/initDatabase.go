package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *sql.DB
var GormDB *gorm.DB

func InitGorm() error {
	// Connect via Gorm
	dsn := "root:M@rk8478@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	GormDB = database

	return err
}

func InitDB() error {
	// Connect via direct MySQL connection
	database, err := sql.Open("mysql", "root:M@rk8478@tcp(127.0.0.1:3306)/store")
	if err != nil {
		return err
	}

	Db = database

	return Db.Ping()
}
