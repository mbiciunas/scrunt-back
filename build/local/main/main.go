package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"scrunt-back/build/local"
	"time"
)

func main() {
	fmt.Println("Delete gorm.db")
	err := os.Remove(".scrunt/database/gorm.db")
	if err != nil {
		fmt.Println("Failed to delete the SQLite database.", err)
	}

	fmt.Println("Open Sqlite database")
	// Open a new connection to our local database.
	db, err := gorm.Open(sqlite.Open(".scrunt/database/gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	// Create the table from our struct.
	err = db.AutoMigrate(
		&local.Script{},
		&local.Service{},
		&local.ServiceType{},
		&local.ScriptServiceType{},
		&local.Version{},
		&local.Code{},
		&local.VersionCode{},
		&local.VersionService{},
		&local.TagType{},
		&local.Tag{},
		&local.ScriptTag{},
		&local.Icon{},
	)

	if err != nil {
		panic("Failed to automigrate.")
	}

	local.InsertTagTypes(db)
	local.InsertTags(db)
	local.InsertIcons(db)
	local.InsertScripts(db)
	local.InsertScriptTags(db)
	local.InsertVersions(db)
	local.InsertCodes(db)
	local.InsertVersionCodes(db)

	// Create a new user in our database.
	db.Create(&local.Script{
		Name:      "Script 1",
		DescLong:  "Long Description 1",
		DescShort: "Short Description 1",
		Created:   time.Time{},
	})

	// Find all of our scripts.
	var scripts []local.Script
	db.Find(&scripts)

	// Output the users from the DB json encoded
	jsonEncoded, _ := json.Marshal(&scripts)
	fmt.Println(string(jsonEncoded))

}
