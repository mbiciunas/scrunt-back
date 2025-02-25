package code

import (
	"scrunt-back/models"
	"scrunt-back/models/scrunt"
)

func GormInsertCode(codetype string, value string, uuid string) (uint, error) {
	code := models.Code{
		Type:  codetype,
		Value: value,
		Uuid:  uuid,
	}

	if err := scrunt.GormDB.Create(&code).Error; err != nil {
		return 0, err
	}

	return code.Id, nil
}
