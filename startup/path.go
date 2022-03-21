package startup

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const SCRUNT = ".scrunt"
const FRONT = "front"
const PYTHON = "python"
const DATABASE = "database"

var pathFront = filepath.Join(SCRUNT, FRONT)
var pathPython = filepath.Join(SCRUNT, PYTHON)
var pathDatabase = filepath.Join(SCRUNT, DATABASE)

func GetPathFront(filename string) string {
	return filepath.Join(pathFront, filename)
}

func GetPathPython(filename string) string {
	return filepath.Join(pathPython, filename)
}

func GetPathDatabase(filename string) string {
	return filepath.Join(pathDatabase, filename)
}

func CurrentDirectory() {
	fmt.Println("Current Directory *******************")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path) // for example /home/user}

	fmt.Println("*************************************")
}

func MakeDirectory() {
	fmt.Println("Make Directory **********************")
	err := os.MkdirAll(SCRUNT, 0700)
	if err != nil {
		log.Println(err)
	}

	err = os.MkdirAll(pathFront, 0700)
	if err != nil {
		log.Println(err)
	}

	err = os.MkdirAll(pathPython, 0600)
	if err != nil {
		log.Println(err)
	}

	err = os.MkdirAll(pathDatabase, 0600)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("*************************************")
}
