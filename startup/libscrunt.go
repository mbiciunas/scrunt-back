package startup

import (
	"io/fs"
)

func libScrunt(fileSystem fs.FS) {
	WriteFilesystem(fileSystem, "libscrunt")
}
