package code

import (
	"github.com/google/uuid"
)

func GenerateCodeUUID(codeType string, value string) string {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(codeType+value)).String()
}
