package startup

import "embed"

func Startup(embeddedFiles embed.FS) {

	makeDirectory()

	frontend(embeddedFiles)

	python(embeddedFiles)

	database(embeddedFiles)
}
