package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"time"
)

type Service struct {
	Id              uint             `json:"id" gorm:"primarykey"`
	IconId          uint             `json:"icon_id"`
	Name            string           `json:"name"`
	Description     string           `json:"description"`
	Address         string           `json:"address"`
	Port            int              `json:"description_long"`
	ServiceTypeId   uint             `json:"service_type_id"`
	VersionServices []VersionService `gorm:"foreignkey:ServiceId"`
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
	IconId             uint                `json:"icon_id"`
	Name               string              `json:"name"`
	DescriptionShort   string              `json:"description_short"`
	DescriptionLong    string              `json:"description_long"`
	Code               string              `json:"code"`
	ScriptServiceTypes []ScriptServiceType `gorm:"foreignkey:ScriptId"`
	Versions           []Version           `gorm:"foreignkey:ScriptId"`
	ScriptTag          []ScriptTag         `gorm:"foreignkey:ScriptId"`
}

type Version struct {
	Id              uint             `json:"id" gorm:"primarykey"`
	ScriptId        uint             `json:"script_id"`
	Created         time.Time        `json:"service_type_id"`
	Major           int64            `json:"major"`
	Minor           int64            `json:"minor"`
	Patch           int              `json:"patch"`
	Commit          int              `json:"commit"`
	Uuid            string           `json:"uuid"`
	Contents        string           `json:"changes"`
	VersionCodes    []VersionCode    `gorm:"foreignkey:VersionId"`
	VersionServices []VersionService `gorm:"foreignkey:VersionId"`
}

type Code struct {
	Id           uint          `json:"id" gorm:"primarykey"`
	Type         string        `json:"type"`
	Value        string        `json:"value"`
	Uuid         string        `json:"uuid"`
	VersionCodes []VersionCode `gorm:"foreignkey:CodeId"`
}

type VersionCode struct {
	Id        uint `json:"id" gorm:"primarykey"`
	VersionId uint `json:"version_id"`
	CodeId    uint `json:"code_id"`
}

type VersionService struct {
	Id         uint   `json:"id" gorm:"primarykey"`
	VersionId  uint   `json:"version_id"`
	ServiceId  uint   `json:"service_id"`
	VersionMin uint   `json:"version_min"`
	VersionMax uint   `json:"version_max"`
	Usage      string `json:"usage"`
	Details    string `json:"details"`
}

type TagType struct {
	Id          uint   `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tag         []Tag  `gorm:"foreignkey:TagTypeId"`
}

type Tag struct {
	Id        uint        `json:"id" gorm:"primarykey"`
	TagTypeId uint        `json:"tag_type_id"`
	Name      string      `json:"name"`
	ScriptTag []ScriptTag `gorm:"foreignkey:TagId"`
}

type ScriptTag struct {
	Id       uint `json:"id" gorm:"primarykey"`
	ScriptId uint `json:"script_id"`
	TagId    uint `json:"tag_id"`
}

type Icon struct {
	Id        uint    `json:"id" gorm:"primarykey"`
	Code      uint    `json:"script_id"`
	Directory string  `json:"directory"`
	Filename  string  `json:"filename"`
	Service   Service `gorm:"foreignkey:IconId"`
	Script    Script  `gorm:"foreignkey:IconId"`
}

func main() {
	fmt.Println("Delete gorm.db")
	err := os.Remove("./gorm.db")
	if err != nil {
		panic("Failed to delete the SQLite database.")
	}

	fmt.Println("Open Sqlite database")
	// Open a new connection to our sqlite database.
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}

	// Create the table from our struct.
	err = db.AutoMigrate(
		&Script{},
		&Service{},
		&ServiceType{},
		&ScriptServiceType{},
		&Version{},
		&Code{},
		&VersionCode{},
		&VersionService{},
		&TagType{},
		&Tag{},
		&ScriptTag{},
		&Icon{},
	)

	if err != nil {
		panic("Failed to automigrate.")
	}

	// Create a new user in our database.
	db.Create(&Script{
		Name:             "Script 1",
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
