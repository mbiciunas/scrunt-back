package local

import (
	"fmt"
	"scrunt-back/models/scrunt/tag"
)

// Tags
var tagIdInstall uint
var tagIdUpgrade uint
var tagIdUninstall uint
var tagIdLogging uint
var tagIdConfiguration uint
var tagIdOther uint
var tagIdBackup uint
var tagIdRestore uint

func InsertTags() (err error) {
	fmt.Println("Insert Tags")

	tagIdInstall, err = tag.GormInsertTag(tagTypeIdPurpose, "Install")
	tagIdUpgrade, err = tag.GormInsertTag(tagTypeIdPurpose, "Upgrade")
	tagIdUninstall, err = tag.GormInsertTag(tagTypeIdPurpose, "Uninstall")
	tagIdLogging, err = tag.GormInsertTag(tagTypeIdPurpose, "Logging")
	tagIdConfiguration, err = tag.GormInsertTag(tagTypeIdPurpose, "Configuration")
	tagIdOther, err = tag.GormInsertTag(tagTypeIdPurpose, "Other")
	tagIdBackup, err = tag.GormInsertTag(tagTypeIdPurpose, "Backup")
	tagIdRestore, err = tag.GormInsertTag(tagTypeIdPurpose, "Restore")

	return err
}
