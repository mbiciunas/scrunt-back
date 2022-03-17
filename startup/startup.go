package startup

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

func ListFiles(embededFiles embed.FS) {
	fmt.Println("*************************************")
	files, err := embededFiles.ReadDir("frontend")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir(), file.Type())
	}

	fmt.Println("*************************************")
}

func ListFilesAll(embededFiles embed.FS) {
	fmt.Println("*************************************")

	if err := run(embededFiles); err != nil {
		log.Fatal(err)
	}

	fmt.Println("*************************************")
}

func run(embededFiles embed.FS) error {
	return fs.WalkDir(embededFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())
		return nil
	})
}
