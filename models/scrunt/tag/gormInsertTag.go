package tag

import (
	"scrunt-back/models/scrunt"
)

func GormInsertTag(tagTypeId uint, name string) (uint, error) {
	tag := scrunt.Tag{
		TagTypeId: tagTypeId,
		Name:      name,
	}

	if err := scrunt.GormDB.Create(&tag).Error; err != nil {
		return 0, err
	}

	return tag.Id, nil
}
