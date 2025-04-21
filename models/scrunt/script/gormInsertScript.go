package script

import (
	"scrunt-back/models/scrunt"
	"time"
)

func GormInsertScript(uuid string, name string, iconCode string, descShort string, descLong string, source scrunt.ScriptSource, parent string, created time.Time) (uint, error) {
	script := scrunt.Script{
		Uuid:      uuid,
		Name:      name,
		IconCode:  iconCode,
		DescShort: descShort,
		DescLong:  descLong,
		Source:    source,
		Parent:    parent,
		Created:   created,
	}

	if err := scrunt.GormDB.Create(&script).Error; err != nil {
		return 0, err
	}

	return script.Id, nil
}
