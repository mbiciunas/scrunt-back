package tagtype

import (
	"scrunt-back/models/scrunt"
)

func GormInsertTagType(name string, description string) (uint, error) {
	tagType := scrunt.TagType{
		Name: name,
		Desc: description,
	}

	err := scrunt.GormDB.Create(&tagType).Error
	if err != nil {
		return 0, err
	}

	return tagType.Id, nil
}
