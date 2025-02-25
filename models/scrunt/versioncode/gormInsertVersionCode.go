package versioncode

import (
	"scrunt-back/models"
	"scrunt-back/models/scrunt"
)

func GormInsertVersionCode(versionId uint, codeId uint, order uint) (uint, error) {
	versionCode := models.VersionCode{
		VersionId: versionId,
		CodeId:    codeId,
		SortOrder: order,
	}

	if err := scrunt.GormDB.Create(&versionCode).Error; err != nil {
		return 0, err
	}

	return versionCode.Id, nil
}
