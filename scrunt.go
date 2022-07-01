package main

import (
	"embed"
	"github.com/pkg/browser"
	"log"
	"scrunt-back/models"
	"scrunt-back/models/runtime"
	"scrunt-back/python"
	"scrunt-back/startup"
)

//go:embed embed
var embeddedFiles embed.FS

func main() {
	// Run startup to extract files from embed and write to .scrunt
	startup.Startup(embeddedFiles)

	// Initialize the router
	router := router()

	// Set up the API endpoints
	api(router)

	// Test python working
	python.Python(1, "print('hello world from simple')")

	// Open browser page
	//openBrowser()

	// Connect to the scrunt database
	err := models.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the runtime database
	err = runtime.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Start and run the server
	err = router.Run(":8080")
	if err != nil {
		return
	}

	//fileServer := http.FileServer(http.Dir("./.scrunt/frontend")) // New code
	//http.Handle("/", fileServer)                                  // New code
	//
	//fmt.Printf("Starting server at port 8080\n")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}
}

func openBrowser() {
	err := browser.OpenURL("http://localhost:8080")
	if err != nil {
		return
	}
}
