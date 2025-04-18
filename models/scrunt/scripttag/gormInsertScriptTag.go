package scripttag

import (
	"scrunt-back/models/scrunt"
)

func GormInsertScriptTag(scriptId uint, tagId uint) (uint, error) {
	scriptTag := scrunt.ScriptTag{
		ScriptId: scriptId,
		TagId:    tagId,
	}

	if err := scrunt.GormDB.Create(&scriptTag).Error; err != nil {
		return 0, err
	}

	return scriptTag.Id, nil
}
