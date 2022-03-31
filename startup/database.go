package startup

import (
	"embed"
)

func database(embeddedFiles embed.FS) {
	WriteFilesystem(embeddedFiles, "embed/database", "database")
}
