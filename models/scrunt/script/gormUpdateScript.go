package script

import (
	_ "modernc.org/sqlite"
	"scrunt-back/models/scrunt"
)

func GormUpdateScript(id uint, name string, iconCode string, descShort string, descLong string) (int64, error) {
	script := scrunt.Script{
		Id:        id,
		Name:      name,
		IconCode:  iconCode,
		DescShort: descShort,
		DescLong:  descLong,
	}

	result := scrunt.GormDB.Updates(&script)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
