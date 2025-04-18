package version

import (
	"github.com/google/uuid"
	"strconv"
)

func GenerateVersionUUID(scriptId uint, major uint, minor uint, patch uint, save uint, change string) string {
	data := strconv.Itoa(int(scriptId)) +
		strconv.Itoa(int(major)) +
		strconv.Itoa(int(minor)) +
		strconv.Itoa(int(patch)) +
		strconv.Itoa(int(save)) +
		change

	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(data)).String()
}
