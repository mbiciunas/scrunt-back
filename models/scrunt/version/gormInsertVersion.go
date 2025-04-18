package version

import (
	"github.com/google/uuid"
	"scrunt-back/models/scrunt"
	"strconv"
	"time"
)

func GormInsertVersion(scriptId uint, time time.Time, major int, minor int, patch int, save int, change string) (uint, error) {
	uuidVersion := genUUID(scriptId, major, minor, patch, save, change)

	version := scrunt.Version{
		ScriptId: scriptId,
		Created:  time.UTC(),
		Major:    major,
		Minor:    minor,
		Patch:    patch,
		Save:     save,
		Uuid:     uuidVersion,
		Change:   change,
	}

	if err := scrunt.GormDB.Create(&version).Error; err != nil {
		return 0, err
	}

	return version.Id, nil
}

func genUUID(scriptId uint, major int, minor int, patch int, save int, change string) string {
	data := strconv.Itoa(int(scriptId)) +
		strconv.Itoa(major) +
		strconv.Itoa(minor) +
		strconv.Itoa(patch) +
		strconv.Itoa(save) +
		change

	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(data)).String()
}
