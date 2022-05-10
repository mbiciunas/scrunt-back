package python

import (
	"fmt"
	"github.com/DataDog/go-python3"
	"scrunt-back/startup"
)

func Python(code string) {
	pythonPath := startup.GetPathPython("./lib/python3.7")
	fmt.Println(pythonPath)

	err := python3.Py_SetPath(pythonPath)
	if err != nil {
		return
	}

	python3path, _ := python3.Py_GetPath()

	fmt.Println("python3.Py_GetPath(): ", python3path)

	python3.Py_Initialize()

	python3.PyRun_SimpleString(code)

	python3.Py_Finalize()
}
