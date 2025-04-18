package icon

import (
	"scrunt-back/models/scrunt"
)

func GormInsertIcon(code string, directory string, filename string) (uint, error) {
	icon := scrunt.Icon{
		Code:      code,
		Directory: directory,
		Filename:  filename,
	}

	if err := scrunt.GormDB.Create(&icon).Error; err != nil {
		return 0, err
	}

	return icon.Id, nil
}
