package main

import (
	"embed"
	"fmt"
	"github.com/DataDog/go-python3"
	"github.com/pkg/browser"
	"log"
	"net/http"
	"scrunt-back/models"
	"scrunt-back/startup"
)

//go:embed embed
var embeddedFiles embed.FS

func main() {
	// Run startup to extract files from embed and write to .scrunt
	startup.Startup(embeddedFiles)

	// Test python working
	python()

	// Open browser page
	//openBrowser()

	err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.Dir("./.scrunt/frontend")) // New code
	http.Handle("/", fileServer)                                  // New code

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

func python() {
	pythonPath := startup.GetPathPython("./lib/python3.7")
	//pythonPath := startup.GetPathPython("./Python-3.7.12/lib/python3.7")
	fmt.Println(pythonPath)

	err := python3.Py_SetPath(pythonPath)
	//err := python3.Py_SetPath("./python-3.7.12/lib/python3.7")
	if err != nil {
		return
	}

	python3path, _ := python3.Py_GetPath()

	fmt.Println("python3.Py_GetPath(): ", python3path)

	python3.Py_Initialize()

	python3.PyRun_SimpleString("print('hello world from simple')")
	python3.PyRun_SimpleString("print('hello world from simple')")

	python3.Py_Finalize()
}
