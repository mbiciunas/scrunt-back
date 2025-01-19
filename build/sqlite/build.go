package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

type Service struct {
	Id            uint   `json:"id" gorm:"primarykey"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Address       string `json:"address"`
	Port          int    `json:"description_long"`
	ServiceTypeId uint   `json:"service_type_id"`
}

type ServiceType struct {
	Id                 uint                `json:"id" gorm:"primarykey"`
	Name               string              `json:"name"`
	Icon               string              `json:"icon"`
	Services           []Service           `gorm:"foreignkey:ServiceTypeId"`
	ScriptServiceTypes []ScriptServiceType `gorm:"foreignkey:ServiceTypeId"`
}

type ScriptServiceType struct {
	Id            uint   `json:"id" gorm:"primarykey"`
	ScriptId      uint   `json:"script_id"`
	ServiceTypeId uint   `json:"service_type_id"`
	Name          string `json:"name"`
}

type Script struct {
	Id                 uint                `json:"id" gorm:"primarykey"`
	Name               string              `json:"name"`
	Description        string              `json:"description"`
	DescriptionShort   string              `json:"description_short"`
	DescriptionLong    string              `json:"description_long"`
	Code               string              `json:"code"`
	ScriptServiceTypes []ScriptServiceType `gorm:"foreignkey:ScriptId"`
	Versions           []Version           `gorm:"foreignkey:ScriptId"`
}

type Version struct {
	Id       uint      `json:"id" gorm:"primarykey"`
	ScriptId uint      `json:"script_id"`
	Created  time.Time `json:"service_type_id"`
	Major    int64     `json:"major"`
	Minor    int64     `json:"minor"`
	Patch    int       `json:"patch"`
	Commit   int       `json:"commit"`
	Uuid     string    `json:"uuid"`
	Contents string    `json:"changes"`
}

func main() {
	fmt.Println("Build Sqlite")
	// Open a new connection to our sqlite database.
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	// Create the table from our struct.
	err = db.AutoMigrate(&Script{}, &Service{}, &ServiceType{}, &ScriptServiceType{}, &Version{})
	if err != nil {
		panic("Failed to automigrate.")
	}

	// Create a new user in our database.
	db.Create(&Script{
		Name:             "Script 1",
		Description:      "Description 1",
		DescriptionLong:  "Long Description 1",
		DescriptionShort: "Short Description 1",
		Code:             "This is my script code...",
	})

	// Find all of our scripts.
	var scripts []Script
	db.Find(&scripts)

	// Output the users from the DB json encoded
	jsonEncoded, _ := json.Marshal(&scripts)
	fmt.Println(string(jsonEncoded))

}
