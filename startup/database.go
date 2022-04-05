package startup

import (
	"io/fs"
)

func database(fileSystem fs.FS) {
	WriteFilesystem(fileSystem, "database")
}
