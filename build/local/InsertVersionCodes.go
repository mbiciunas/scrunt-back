package local

import (
	"fmt"
	"gorm.io/gorm"
)

// ScriptTags
var versionCode1 VersionCode
var versionCode2 VersionCode
var versionCode3 VersionCode

func InsertVersionCodes(db *gorm.DB) {
	fmt.Println("Insert Versions")

	versionCode1 = VersionCode{VersionId: scriptEC2.Id, CodeId: codeImport.Id, SortOrder: 1}
	versionCode2 = VersionCode{VersionId: scriptEC2.Id, CodeId: codeParameter.Id, SortOrder: 2}
	versionCode3 = VersionCode{VersionId: scriptEC2.Id, CodeId: codeMain.Id, SortOrder: 3}

	db.Create(&versionCode1)
	db.Create(&versionCode2)
	db.Create(&versionCode3)
}
