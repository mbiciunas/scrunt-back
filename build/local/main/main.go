package main

import (
	"fmt"
	"log"
	"os"
	"scrunt-back/build/local"
	"scrunt-back/models/scrunt"
)

func main() {
	fmt.Println("Delete gorm.db")
	err := os.Remove(".scrunt/database/gorm.db")
	if err != nil {
		fmt.Println("Failed to delete the SQLite database.", err)
	}

	// Connect to the scrunt database with Gorm
	err = scrunt.InitGorm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Auto-migrate to create the tables")
	err = scrunt.GormDB.AutoMigrate(
		&scrunt.Script{},
		&scrunt.Service{},
		&scrunt.ServiceType{},
		&scrunt.ScriptServiceType{},
		&scrunt.Version{},
		&scrunt.Code{},
		&scrunt.VersionCode{},
		&scrunt.VersionService{},
		&scrunt.TagType{},
		&scrunt.Tag{},
		&scrunt.ScriptTag{},
		&scrunt.Icon{},
	)

	if err != nil {
		panic("Failed to auto-migrate.")
	}

	err = local.InsertTagTypes()
	if err != nil {
		return
	}

	err = local.InsertTags()
	if err != nil {
		return
	}

	err = local.InsertIcons()
	if err != nil {
		return
	}

	err = local.InsertScripts()
	if err != nil {
		return
	}

	err = local.InsertScriptTags()
	if err != nil {
		return
	}

	err = local.InsertVersions()
	if err != nil {
		return
	}

	err = local.InsertCodes()
	if err != nil {
		return
	}

	err = local.InsertVersionCodes()
	if err != nil {
		return
	}
}
