package main

import (
	"embed"
	"fmt"
	"scrunt-back/models/scrunt/icon"
	"scrunt-back/models/scrunt/script"
	"scrunt-back/models/scrunt/scripttag"
	"scrunt-back/models/scrunt/tag"
	"scrunt-back/models/scrunt/tagtype"
	"scrunt-back/models/scrunt/version"
	"time"

	//"github.com/pkg/browser"
	"log"
	"scrunt-back/models/runtime"
	"scrunt-back/models/scrunt"
	"scrunt-back/models/store"
	"scrunt-back/startup"
)

//go:embed embed
var embeddedFiles embed.FS

func main() {
	// Run startup to extract files from embed and write to .scrunt
	if startup.InstallRequired() {
		startup.Startup(embeddedFiles)
	}

	// Initialize the router
	router := router()

	// Set up the API endpoints
	apiScrunt(router)
	apiStore(router)

	// Open browser page
	//openBrowser()

	// Connect to the store database with Gorm
	err := store.InitGorm()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the store database
	err = store.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the scrunt database with Gorm
	err = scrunt.InitGorm()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the scrunt database
	err = scrunt.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the runtime database
	err = runtime.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	// THIS IS TEMPORARY!!!  TESTING ADDING TO SQLITE USING GORM...
	// Create a new script in our database.
	tagTypeId, err := tagtype.GormInsertTagType("My tag type", "This is a general tag type")
	fmt.Println("Inserted a tag type", tagTypeId)
	tagId, err := tag.GormInsertTag(tagTypeId, "My Tag")
	fmt.Println("Inserted a tag", tagId)
	iconIdAws, err := icon.GormInsertIcon("aws.aws", "aws", "aws.svg")
	fmt.Println("Inserted a icon", iconIdAws)
	scriptId, err := script.GormInsertScript(iconIdAws, "Script 1", "Short Description 1", "Long Description 1", time.Now().UTC())
	fmt.Println("Inserted a script", scriptId)
	scriptTagId, err := scripttag.GormInsertScriptTag(scriptId, tagId)
	fmt.Println("Inserted a scripttag", scriptTagId)
	versionId, err := version.GormInsertVersion(scriptId, time.Now(), 1, 0, 0, 0, "Created new version")
	fmt.Println("Inserted a version", versionId)

	// Initialize the python instance.
	// Note we're not deferring finalize here since we need the instance
	// for as long as scrunt is running.
	//python.Initialize()

	// Start and run the server
	err = router.Run(":8080")
	if err != nil {
		return
	}

	// Finalize python since we're exiting the program now.
	//python.Finalize()

	//fileServer := http.FileServer(http.Dir("./.scrunt/frontend")) // New code
	//http.Handle("/", fileServer)                                  // New code
	//
	//fmt.Printf("Starting server at port 8080\n")
	//if err := http.ListenAndServe(":8080", nil); err != nil {
	//	log.Fatal(err)
	//}
}

//func openBrowser() {
//	err := browser.OpenURL("http://localhost:8080")
//	if err != nil {
//		return
//	}
//}
