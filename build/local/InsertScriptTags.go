package local

import (
	"fmt"
	"gorm.io/gorm"
)

// ScriptTags
var scriptTagEc2Install ScriptTag
var scriptTagMysqlBackup ScriptTag
var scriptTagMysqlRestore ScriptTag
var scriptTagChatgptOther ScriptTag

func InsertScriptTags(db *gorm.DB) {
	fmt.Println("Insert Script Tags")

	scriptTagEc2Install = ScriptTag{ScriptId: scriptEC2.Id, TagId: tagInstall.Id}
	scriptTagMysqlBackup = ScriptTag{ScriptId: scriptMySQL.Id, TagId: tagBackup.Id}
	scriptTagMysqlRestore = ScriptTag{ScriptId: scriptMySQL.Id, TagId: tagRestore.Id}
	scriptTagChatgptOther = ScriptTag{ScriptId: scriptChatGPT.Id, TagId: tagOther.Id}

	db.Create(&scriptTagEc2Install)
	db.Create(&scriptTagMysqlBackup)
	db.Create(&scriptTagMysqlRestore)
	db.Create(&scriptTagChatgptOther)
}
