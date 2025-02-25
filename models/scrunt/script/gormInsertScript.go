package script

import (
	"scrunt-back/models"
	"scrunt-back/models/scrunt"
	"time"
)

func GormInsertScript(iconId uint, name string, descShort string, descLong string, created time.Time) (uint, error) {
	script := models.Script{
		IconId:    iconId,
		Name:      name,
		DescShort: descShort,
		DescLong:  descLong,
		Created:   created,
	}

	if err := scrunt.GormDB.Create(&script).Error; err != nil {
		return 0, err
	}

	return script.Id, nil
}
