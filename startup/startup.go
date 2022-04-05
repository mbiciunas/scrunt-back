package startup

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"log"
)

func Startup(embeddedFiles embed.FS) *gin.Engine {

	fileSystem, err := fs.Sub(fs.FS(embeddedFiles), "embed")
	if err != nil {
		log.Fatal(err)
	}

	makeDirectory()

	frontend(fileSystem)

	python(fileSystem)

	database(fileSystem)

	router := router()

	api(router)

	return router
}
