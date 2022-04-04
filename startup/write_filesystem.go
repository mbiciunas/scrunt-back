package startup

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func WriteFilesystem(fileSystem fs.FS, pathRead string) {
	fmt.Println("Write Files *************************")

	err := fs.WalkDir(fileSystem, pathRead, func(pathEmbed string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Printf("pathEmbed=%q, isDir=%v\n", pathEmbed, d.IsDir())

		if d.IsDir() {
			err = os.MkdirAll(filepath.Join(pathScrunt, pathEmbed), 0700)
			if err != nil {
				log.Println(err)
			}
		} else {
			writeFile(fileSystem, pathEmbed)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*************************************")
}

func writeFile(fileSystem fs.FS, pathEmbed string) {
	// Open file for reading
	fileRead, err := fileSystem.Open(pathEmbed)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = fileRead.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Create file for writing
	fileWrite, err := os.Create(filepath.Join(pathScrunt, pathEmbed))
	if err != nil {
		fmt.Println("Did not work writing to:", filepath.Join(pathScrunt, pathEmbed))
		log.Fatal(err)
	}
	defer func() {
		if err = fileWrite.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Set up the buffers
	bufferRead := bufio.NewReader(fileRead)
	bufferWriter := bufio.NewWriter(fileWrite)
	buffer := make([]byte, 4096)

	for {
		// Read a chunk of the source file
		n, err := bufferRead.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal("Error reading file:", err)
		}
		if n == 0 {
			break
		}

		// Write a chunk to the destination file
		_, err = bufferWriter.Write(buffer[:n])
		if err != nil {
			log.Fatal(err)
		}
	}

	// Flush the writes to disk
	err = bufferWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
