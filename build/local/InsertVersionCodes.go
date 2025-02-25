package local

import (
	"fmt"
	"scrunt-back/models/scrunt/versioncode"
)

// ScriptTags
var versionCodeId1 uint
var versionCodeId2 uint
var versionCodeId3 uint

func InsertVersionCodes() (err error) {
	fmt.Println("Insert Version Codes")

	versionCodeId1, err = versioncode.GormInsertVersionCode(scriptIdEC2, codeIdImport, 1)
	versionCodeId2, err = versioncode.GormInsertVersionCode(scriptIdEC2, codeIdParameter, 2)
	versionCodeId3, err = versioncode.GormInsertVersionCode(scriptIdEC2, codeIdMain, 3)

	return err
}
