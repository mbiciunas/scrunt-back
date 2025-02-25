package local

import (
	"fmt"
	"scrunt-back/models/scrunt/scripttag"
)

// ScriptTags
var scriptTagIdEc2Install uint
var scriptTagIdMysqlBackup uint
var scriptTagIdMysqlRestore uint
var scriptTagIdChatgptOther uint

func InsertScriptTags() (err error) {
	fmt.Println("Insert Script Tags")

	scriptTagIdEc2Install, err = scripttag.GormInsertScriptTag(scriptIdEC2, tagIdInstall)
	scriptTagIdMysqlBackup, err = scripttag.GormInsertScriptTag(scriptIdMySQL, tagIdBackup)
	scriptTagIdMysqlRestore, err = scripttag.GormInsertScriptTag(scriptIdMySQL, tagIdRestore)
	scriptTagIdChatgptOther, err = scripttag.GormInsertScriptTag(scriptIdChatGPT, tagIdOther)

	return err
}
