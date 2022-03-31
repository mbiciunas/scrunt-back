package startup

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func WriteFilesystem(embeddedFiles embed.FS, path string) {
	fmt.Println("List All Files **********************")

	err := fs.WalkDir(embeddedFiles, path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())

		if d.IsDir() {
			err = os.MkdirAll(filepath.Join(pathScrunt, path), 0700)
			if err != nil {
				log.Println(err)
			}
		} else {
			writeFile(embeddedFiles, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*************************************")
}

func writeFile(embeddedFiles embed.FS, path string) {
	//fmt.Println("Write File **************************")

	file, err := embeddedFiles.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Path: " + path + "  Directory: " + GetPathScrunt(path))
	err = ioutil.WriteFile(GetPathScrunt(path), file, 0644)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("*************************************")
}
