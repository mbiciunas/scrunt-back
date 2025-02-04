package local

import (
	"fmt"
	"gorm.io/gorm"
)

// Tags
var tagInstall Tag
var tagUpgrade Tag
var tagUninstall Tag
var tagLogging Tag
var tagConfiguration Tag
var tagOther Tag
var tagBackup Tag
var tagRestore Tag

func InsertTags(db *gorm.DB) {
	fmt.Println("Insert Tags")

	tagInstall = Tag{TagTypeId: tagTypePurpose.Id, Name: "Install"}
	tagUpgrade = Tag{TagTypeId: tagTypePurpose.Id, Name: "Upgrade"}
	tagUninstall = Tag{TagTypeId: tagTypePurpose.Id, Name: "Uninstall"}
	tagLogging = Tag{TagTypeId: tagTypePurpose.Id, Name: "Logging"}
	tagConfiguration = Tag{TagTypeId: tagTypePurpose.Id, Name: "Configuration"}
	tagOther = Tag{TagTypeId: tagTypePurpose.Id, Name: "Other"}
	tagBackup = Tag{TagTypeId: tagTypePurpose.Id, Name: "Backup"}
	tagRestore = Tag{TagTypeId: tagTypePurpose.Id, Name: "Restore"}

	db.Create(&tagInstall)
	db.Create(&tagUpgrade)
	db.Create(&tagUninstall)
	db.Create(&tagLogging)
	db.Create(&tagConfiguration)
	db.Create(&tagOther)
	db.Create(&tagBackup)
	db.Create(&tagRestore)
}
