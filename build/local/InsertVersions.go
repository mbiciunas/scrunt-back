package local

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// ScriptTags
var versionScriptEC2100 Version
var versionScriptEC2101 Version
var versionScriptEC2102 Version
var versionScriptEC2110 Version
var versionScriptEC2120 Version
var versionScriptEC2121 Version
var versionScriptEC2122 Version
var versionScriptEC2200 Version

func InsertVersions(db *gorm.DB) {
	fmt.Println("Insert Versions")

	versionScriptEC2100 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 1, 31, 13, 46, 23, 0, time.Local), Major: 1, Minor: 0, Patch: 0, Commit: 0, Uuid: genUUID("1000"), Change: "Script created"}
	versionScriptEC2101 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 2, 21, 16, 23, 40, 0, time.Local), Major: 1, Minor: 0, Patch: 1, Commit: 0, Uuid: genUUID("1010"), Change: "Fixed broken call"}
	versionScriptEC2102 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 3, 02, 16, 02, 13, 0, time.Local), Major: 1, Minor: 0, Patch: 2, Commit: 0, Uuid: genUUID("1020"), Change: "Added error checking"}
	versionScriptEC2110 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 6, 17, 9, 12, 03, 0, time.Local), Major: 1, Minor: 1, Patch: 0, Commit: 0, Uuid: genUUID("1100"), Change: "Added support for selecting a region"}
	versionScriptEC2120 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 8, 24, 9, 22, 32, 0, time.Local), Major: 1, Minor: 2, Patch: 0, Commit: 0, Uuid: genUUID("1200"), Change: "Added support for selecting a machine size"}
	versionScriptEC2121 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 8, 25, 06, 34, 54, 0, time.Local), Major: 1, Minor: 2, Patch: 1, Commit: 0, Uuid: genUUID("1210"), Change: "Corrected call to AWS"}
	versionScriptEC2122 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 8, 27, 11, 43, 12, 0, time.Local), Major: 1, Minor: 2, Patch: 2, Commit: 0, Uuid: genUUID("1220"), Change: "Added error checking on the machine size"}
	versionScriptEC2200 = Version{ScriptId: scriptEC2.Id, Created: time.Date(2021, 9, 9, 9, 12, 9, 0, time.Local), Major: 2, Minor: 0, Patch: 0, Commit: 0, Uuid: genUUID("2000"), Change: "Added operating system, disk size support"}

	db.Create(&versionScriptEC2100)
	db.Create(&versionScriptEC2101)
	db.Create(&versionScriptEC2102)
	db.Create(&versionScriptEC2110)
	db.Create(&versionScriptEC2120)
	db.Create(&versionScriptEC2121)
	db.Create(&versionScriptEC2122)
	db.Create(&versionScriptEC2200)
}

func genUUID(data string) string {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(data)).String()
}
