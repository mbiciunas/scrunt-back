package script

import (
	"github.com/google/uuid"
	"time"
)

func GenerateScriptUUID(name string, created time.Time) string {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(name+created.String())).String()
}
