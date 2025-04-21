package local

import (
	"fmt"
	"scrunt-back/models/scrunt/version"
	"time"
)

// ScriptTags
var versionIdScriptEC2100 uint
var versionIdScriptEC2101 uint
var versionIdScriptEC2102 uint
var versionIdScriptEC2110 uint
var versionIdScriptEC2120 uint
var versionIdScriptEC2121 uint
var versionIdScriptEC2122 uint
var versionIdScriptEC2200 uint

func InsertVersions() (err error) {
	fmt.Println("Insert Versions")

	versionIdScriptEC2100 = insertVersion(scriptIdEC2, time.Date(2021, 1, 31, 13, 46, 23, 0, time.Local), 1, 0, 0, 0, "Script created")
	versionIdScriptEC2101 = insertVersion(scriptIdEC2, time.Date(2021, 2, 21, 16, 23, 40, 0, time.Local), 1, 0, 1, 0, "Fixed broken call")
	versionIdScriptEC2102 = insertVersion(scriptIdEC2, time.Date(2021, 3, 02, 16, 02, 13, 0, time.Local), 1, 0, 2, 0, "Added error checking")
	versionIdScriptEC2110 = insertVersion(scriptIdEC2, time.Date(2021, 6, 17, 9, 12, 03, 0, time.Local), 1, 1, 0, 0, "Added support for selecting a region")
	versionIdScriptEC2120 = insertVersion(scriptIdEC2, time.Date(2021, 8, 24, 9, 22, 32, 0, time.Local), 1, 2, 0, 0, "Added support for selecting a machine size")
	versionIdScriptEC2121 = insertVersion(scriptIdEC2, time.Date(2021, 8, 25, 06, 34, 54, 0, time.Local), 1, 2, 1, 0, "Corrected call to AWS")
	versionIdScriptEC2122 = insertVersion(scriptIdEC2, time.Date(2021, 8, 27, 11, 43, 12, 0, time.Local), 1, 2, 2, 0, "Added error checking on the machine size")
	versionIdScriptEC2200 = insertVersion(scriptIdEC2, time.Date(2021, 9, 9, 9, 12, 9, 0, time.Local), 2, 0, 0, 0, "Added operating system, disk size support")

	return err
}

func insertVersion(scriptId uint, created time.Time, major uint, minor uint, patch uint, save uint, change string) uint {
	uuid := version.GenerateVersionUUID(scriptId, major, minor, patch, save, change)
	versionId, err := version.GormInsertVersion(uuid, scriptId, created, major, minor, patch, save, change)

	if err != nil {
		panic(err)
	}

	return versionId
}
