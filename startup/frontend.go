package startup

import (
	"embed"
)

func frontend(embeddedFiles embed.FS) {
	WriteFilesystem(embeddedFiles, "embed/frontend", "frontend")
}
