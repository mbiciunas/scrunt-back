package main

import (
	"embed"
	"fmt"
	"github.com/pkg/browser"
	"io/fs"
	"log"
	"net/http"
	"scrunt-back/startup"
)

//go:embed frontend
var embededFiles embed.FS

func main() {
	startup.ListFiles(embededFiles)

	startup.ListFilesAll(embededFiles)

	startup.OpenFile(embededFiles)

	//openBrowser()

	fileServer := http.FileServer(getFileSystem()) // New code
	http.Handle("/", fileServer)                   // New code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func openBrowser() {
	err := browser.OpenURL("http://localhost:8080")
	if err != nil {
		return
	}
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "frontend")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}
