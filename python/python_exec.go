package python

import (
	"fmt"
	"os"
	"path/filepath"

	"log"
	"os/exec"
	"scrunt-back/startup"
	"strconv"
)

func Exec(runId int, code string) {
	//
	// Create runtime directory
	//
	createRuntimeDir()

	//
	// Create Run Directory
	//
	runDirectory := createRunDir()

	//defer os.RemoveAll(runDirectory)

	//
	// Create python program to run script
	//
	runFile := createPythonRunner(runDirectory, runId)

	//
	// Save script file in run directory
	//
	createPythonScript(runDirectory, code)

	//
	// Save script file in run directory
	//
	createPythonPackage(runDirectory)

	cmd := exec.Command(".scrunt/python/bin/python3", runFile)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("exec.Command() failed with %s\n\n%s", err, out)
	}
}

func createRuntimeDir() {
	err := os.MkdirAll(startup.GetDirectoryRuntime(), os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Current runtime directory:", startup.GetDirectoryRuntime())
}

func createRunDir() string {
	runDir, err := os.MkdirTemp(startup.GetDirectoryRuntime(), "run-*")

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Run directory:", runDir)

	return runDir
}

func createPythonRunner(runDirectory string, runId int) string {
	scriptFile := filepath.Join(runDirectory, "runner.py")

	//
	// Create Runner file in run directory
	//
	f, err := os.Create(scriptFile)
	if err != nil {
		log.Fatalf("Failed to create runner file: %s", err)
	}
	fmt.Println("Runner file name:", f.Name())

	script := "from libscrunt.scrunt import Scrunt\n"
	script += "import runner_logging\n"
	script += "\n"
	script += "_scrunt = Scrunt(" + strconv.Itoa(runId) + ")\n"
	script += "\n"
	script += "try:\n"
	script += "    runner_logging.main(_scrunt)\n"
	script += "except SystemExit as systemExit:\n"
	script += "    _scrunt.log().error(\"System exit - {}\".format(systemExit.args))\n"
	script += "except KeyboardInterrupt as keyboardInterrupt:\n"
	script += "    _scrunt.log().error(\"Keyboard Interrupt - {}\".format(keyboardInterrupt.args))\n"
	script += "except Exception as exception:\n"
	script += "    _scrunt.log().critical(\"This is a critical exception\")\n"

	_, err = f.Write([]byte(script))
	if err != nil {
		log.Fatalf("Failed to write runner file: %s", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("Failed to close runner file: %s", err)
	}

	return scriptFile
}

func createPythonScript(runDirectory string, code string) string {
	scriptFile := filepath.Join(runDirectory, "runner_logging.py")

	//
	// Create script file in run directory
	//
	f, err := os.Create(scriptFile)
	if err != nil {
		log.Fatalf("Failed to create script file: %s", err)
	}
	fmt.Println("Script file name:", f.Name())

	_, err = f.Write([]byte(code))
	if err != nil {
		log.Fatalf("Failed to write to script file: %s", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("Failed to close script file: %s", err)
	}

	return scriptFile
}

func createPythonPackage(runDirectory string) {
	packageFile := filepath.Join(runDirectory, "__init__.py")

	//
	// Create package file in run directory
	//
	f, err := os.Create(packageFile)
	if err != nil {
		log.Fatalf("Failed to create package file: %s", err)
	}
	fmt.Println("Package file name:", f.Name())

	err = f.Close()
	if err != nil {
		log.Fatalf("Failed to close package file: %s", err)
	}
}
