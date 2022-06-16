package python

import (
	"fmt"
	"github.com/DataDog/go-python3"
	"scrunt-back/python/library"
	"scrunt-back/startup"
)

func Python(code string) {
	library.Write()

	pathPython := startup.GetPathPython("lib/python3.7")
	pathLibDynload := startup.GetPathPython("lib/python3.7/lib-dynload")

	//scriptStart := "import sys\nsys.stdout = open('Output.txt', 'w')\n"
	//scriptEnd := "sys.stdout.close()"
	//script := scriptStart + "\n" + code + "\n" + scriptEnd
	//fmt.Println("Script: ", script)
	fmt.Println("Script: ", code)

	//err := python3.Py_SetPath(pythonPath)
	err := python3.Py_SetPath(pathPython + ":" + pathLibDynload + ":" + "/home/mbiciunas/go/src/scrunt-back")
	if err != nil {
		return
	}

	python3path, _ := python3.Py_GetPythonHome()
	fmt.Println("python3.Py_GetPythonHome(): ", python3path)

	python3.Py_Initialize()
	//fmt.Println("PySys_GetObject", python3.PySys_GetObject("path"))
	fmt.Println("Before Run")
	output := python3.PyRun_SimpleString(code)
	fmt.Println("After Run", output)

	//output = python3.PyRun_SimpleString(script)
	//fmt.Println("After Run", output)

	python3.Py_Finalize()
}
