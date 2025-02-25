package local

import (
	"fmt"
	"scrunt-back/models/scrunt/tagtype"
)

var tagTypeIdPurpose uint

func InsertTagTypes() (err error) {
	fmt.Println("Insert Tag Types")

	tagTypeIdPurpose, err = tagtype.GormInsertTagType("Purpose", "The purpose of the script")
	fmt.Println("Inserted a tag type", tagTypeIdPurpose)

	return err
}
