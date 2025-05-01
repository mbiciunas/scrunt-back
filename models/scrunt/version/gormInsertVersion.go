package version

import (
	"scrunt-back/models/scrunt"
	"time"
)

func GormInsertVersion(uuid string, scriptId uint, created time.Time, major uint, minor uint, patch uint, save uint, change string) (uint, error) {
	version := scrunt.Version{
		Uuid:     uuid,
		ScriptId: scriptId,
		Created:  created.UTC(),
		Major:    major,
		Minor:    minor,
		Patch:    patch,
		Save:     save,
		Change:   change,
	}

	if err := scrunt.GormDB.Create(&version).Error; err != nil {
		return 0, err
	}

	return version.Id, nil
}
