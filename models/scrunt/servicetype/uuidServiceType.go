package servicetype

import (
	"github.com/google/uuid"
)

func GenerateServiceTypeUUID(name string, descShort string) string {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(name+descShort)).String()
}
