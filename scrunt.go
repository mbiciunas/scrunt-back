package main

import (
	"embed"
	"fmt"
	"github.com/pkg/browser"
	"io/fs"
	"log"
	"net/http"
)

//go:embed frontend
var embededFiles embed.FS

func main() {
	fmt.Println("*************************************")
	files, err := embededFiles.ReadDir("frontend")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir(), file.Type())
	}

	fmt.Println("*************************************")

	if err := run(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("*************************************")

	err = browser.OpenURL("http://localhost:8080")
	if err != nil {
		return
	}

	fileServer := http.FileServer(getFileSystem()) // New code
	http.Handle("/", fileServer)                   // New code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func getFileSystem() http.FileSystem {
	fsys, err := fs.Sub(embededFiles, "frontend")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

func run() error {
	return fs.WalkDir(embededFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		return nil
	})
}
