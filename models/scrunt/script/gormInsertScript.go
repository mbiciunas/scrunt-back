package script

import (
	"scrunt-back/models"
	"scrunt-back/models/scrunt"
	"time"
)

func GormInsertScript(name string, iconCode string, descShort string, descLong string, created time.Time) (uint, error) {
	script := models.Script{
		Name:      name,
		IconCode:  iconCode,
		DescShort: descShort,
		DescLong:  descLong,
		Created:   created,
	}

	if err := scrunt.GormDB.Create(&script).Error; err != nil {
		return 0, err
	}

	return script.Id, nil
}
