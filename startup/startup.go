package startup

import (
	"embed"
	"io/fs"
	"log"
	"os"
)

func Startup(embeddedFiles embed.FS) {

	fileSystem, err := fs.Sub(fs.FS(embeddedFiles), "embed")
	if err != nil {
		log.Fatal(err)
	}

	makeDirectory()

	frontend(fileSystem)

	python(fileSystem)

	database(fileSystem)

	libScrunt(fileSystem)

	//return router
}

func InstallRequired() bool {
	var install bool

	if !directoryExist() {
		install = true
	}

	return install
}

func directoryExist() bool {
	_, err := os.Stat(GetDirectoryScrunt())
	if os.IsNotExist(err) {
		return false
	}

	return true
}
