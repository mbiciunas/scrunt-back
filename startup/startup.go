package startup

import (
	"embed"
	"io/fs"
	"log"
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

	//return router
}
