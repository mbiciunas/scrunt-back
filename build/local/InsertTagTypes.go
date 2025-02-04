package local

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

// Tag Types
var tagTypePurpose TagType

func InsertTagTypes(db *gorm.DB) {
	fmt.Println("Insert Tag Types")

	tagTypePurpose = TagType{Name: "Purpose", Desc: "The purpose of the script"}

	db.Create(&tagTypePurpose)
}

func InsertOther(db *gorm.DB) {
	fmt.Println("Insert Tag Types")

	// Create a new user in our database.

	db.Create(&tagTypePurpose)
	fmt.Println(tagTypePurpose.Id)

	// Find all of our scripts.
	var tags []Tag
	db.Find(&tags)

	// Output the users from the DB json encoded
	jsonEncoded, _ := json.Marshal(&tags)
	fmt.Println(string(jsonEncoded))

}
