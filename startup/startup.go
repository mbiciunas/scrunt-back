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

func ListFilesAll(embeddedFiles embed.FS) {
	fmt.Println("List All Files **********************")

	//fileSystem := getFileSystem(embeddedFiles)

	//
	// Trying to convert from embedded file system to regular one without the leading "frontend" part of path.
	// But then can't read the file properly (can't find it).
	// Need to pass fileSystem above into WalkDir and get it opening and reading the contents of the file.
	//
	err := fs.WalkDir(embeddedFiles, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Printf("path=%q, isDir=%v\n", path, d.IsDir())

		if d.IsDir() {
			err = os.MkdirAll(filepath.Join(pathFront, path), 0700)
			if err != nil {
				log.Println(err)
			}
		} else {
			WriteFile(embeddedFiles, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*************************************")
}

func WriteFile(embeddedFiles embed.FS, path string) {
	fmt.Println("Write File **************************")

	file, err := embeddedFiles.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(GetPathFront(path), file, 0644)
	//	err = ioutil.WriteFile("data.txt", file, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*************************************")
}

//func WriteFile(embeddedFiles embed.FS, path string) {
//	fmt.Println("Write File **************************")
//
//	file, err := embeddedFiles.ReadFile(path)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	err = ioutil.WriteFile(GetPathFront(path), file, 0644)
//	//	err = ioutil.WriteFile("data.txt", file, 0644)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("*************************************")
//}

func getFileSystem(embeddedFiles embed.FS) fs.FS {
	fileSystem, err := fs.Sub(embeddedFiles, "frontend")
	if err != nil {
		panic(err)
	}

	return fileSystem
}
