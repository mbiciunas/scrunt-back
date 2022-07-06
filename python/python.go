package python

import (
	"fmt"
	"github.com/DataDog/go-python3"
	"scrunt-back/startup"
	"strconv"
)

func Python(id int, code string) {
	//library.Write()

	//
	// Need to pass the id (should now be the run_id) to the python program.  Best would be as variable or
	// command line parameter.
	// Then need to update libscrunt to use the run_id passed in.
	//

	pathPython := startup.GetPathPython("lib/python3.7")
	pathLibDynload := startup.GetPathPython("lib/python3.7/lib-dynload")
	pathLibScrunt := startup.GetDirectoryLibScrunt()

	//scriptStart := ""
	scriptStart := "from scrunt import Scrunt\nscrunt = Scrunt(" + strconv.Itoa(id) + ")\n\n"
	//scriptStart := "import sys\nsys.stdout = open('Output.txt', 'w')\n"
	//scriptEnd := "sys.stdout.close()"
	script := scriptStart + "\n" + code
	fmt.Println("Script: ", script)
	//fmt.Println("Script: ", code)

	err := python3.Py_SetPath(pathLibScrunt + ":" + pathPython + ":" + pathLibDynload + ":" + "/home/mbiciunas/go/src/scrunt-back")
	if err != nil {
		return
	}

	//python3path, _ := python3.Py_GetPythonHome()
	//fmt.Println("python3.Py_GetPythonHome(): ", python3path)

	python3.Py_Initialize()

	defer python3.Py_Finalize()

	err = python3.PySys_SetArgv([]string{strconv.Itoa(id)})
	if err != nil {
		return
	}

	//
	// Clean up reading of stdOut, stdErr.  There's a link saved with info on how to do.
	// Part of this is writing streams to database so we can see the output.
	//
	// Need to convert to run in libScrunt.  Lots to do here:
	//    - Use standard service from database
	//    - Use credentials from database
	//
	//fmt.Println("PySys_GetObject", python3.PySys_GetObject("path"))
	fmt.Println("Before Run")
	output := python3.PyRun_SimpleString(script)
	fmt.Println("After Run", output)

	//output = python3.PyRun_SimpleString(script)
	//fmt.Println("After Run", output)

	//python3.Py_Finalize()
}
