package version

import (
	"scrunt-back/models/scrunt"
	"time"
)

func GormInsertVersion(uuid string, scriptId uint, time time.Time, major uint, minor uint, patch uint, save uint, change string) (uint, error) {
	//uuidVersion := genUUID(scriptId, major, minor, patch, save, change)

	version := scrunt.Version{
		Uuid:     uuid,
		ScriptId: scriptId,
		Created:  time.UTC(),
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

//func genUUID(scriptId uint, major int, minor int, patch int, save int, change string) string {
//	data := strconv.Itoa(int(scriptId)) +
//		strconv.Itoa(major) +
//		strconv.Itoa(minor) +
//		strconv.Itoa(patch) +
//		strconv.Itoa(save) +
//		change
//
//	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(data)).String()
//}
