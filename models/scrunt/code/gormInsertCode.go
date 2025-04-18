package code

import (
	"scrunt-back/models/scrunt"
)

func GormInsertCode(uuid string, codeType string, value string) (uint, error) {
	code := scrunt.Code{
		Type:  codeType,
		Value: value,
		Uuid:  uuid,
	}

	if err := scrunt.GormDB.Create(&code).Error; err != nil {
		return 0, err
	}

	return code.Id, nil
}
