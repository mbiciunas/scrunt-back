package startup

import (
	"io/fs"
)

func frontend(fileSystem fs.FS) {
	WriteFilesystem(fileSystem, "frontend")
}
