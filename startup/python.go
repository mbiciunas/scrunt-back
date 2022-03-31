package startup

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

func python(embeddedFiles embed.FS) {
	WriteFilesystem(embeddedFiles, "python")

	extractPython()
}

func extractPython() {
	pythonPath := filepath.Join(pathPython, "python.tgz")
	fmt.Println("Current Directory *******************")
	fmt.Println("Python Path: ", pythonPath)
	fmt.Println("Current Directory *******************")
	r, err := os.Open(pythonPath)
	if err != nil {
		fmt.Println("error")
	}

	extractTarGz(pathPython, r)
}
