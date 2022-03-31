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

//
// Need to get rid of leading embed portion of directory.  What we're doing
// now doesn't work with deeper structures.
// May make sense to go back to using SubFS and passing that.  Then fixing the file write stuff.
//

func WriteFilesystem(embeddedFiles embed.FS, pathRead string, pathWrite string) {
	fmt.Println("List All Files **********************")

	err := fs.WalkDir(embeddedFiles, pathRead, func(pathEmbed string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fmt.Printf("pathEmbed=%q, isDir=%v\n", pathEmbed, d.IsDir())

		if d.IsDir() {
			fmt.Println("Create directory:", filepath.Join(pathScrunt, pathWrite))
			err = os.MkdirAll(filepath.Join(pathScrunt, pathWrite), 0700)
			if err != nil {
				log.Println(err)
			}
		} else {
			writeFile(embeddedFiles, pathEmbed, pathWrite)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("*************************************")
}

func writeFile(embeddedFiles embed.FS, pathEmbed string, pathWrite string) {
	file, err := embeddedFiles.ReadFile(pathEmbed)
	if err != nil {
		fmt.Println(err)
	}

	//pathWrite := strings.Replace(pathRead, "embed/", "", 1)

	fmt.Println("Write file:", filepath.Join(pathScrunt, pathWrite))
	err = ioutil.WriteFile(GetPathScrunt(pathWrite), file, 0644)

	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("*************************************")
}
