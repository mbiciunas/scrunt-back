package python

import (
	"errors"
	"github.com/go-python/cpy3"
	"scrunt-back/startup"
)

func Initialize() {
	pathPython := startup.GetPathPython("lib/python3.7")
	pathLibDynload := startup.GetPathPython("lib/python3.7/lib-dynload")
	pathLibScrunt := startup.GetDirectoryLibScrunt()

	err := python3.Py_SetPath(pathLibScrunt + ":" + pathPython + ":" + pathLibDynload + ":" + "/home/mbiciunas/go/src/scrunt-back")
	if err != nil {
		return
	}

	python3.Py_Initialize()

	if !python3.Py_IsInitialized() {
		err = errors.New("error initializing the python interpreter")
		return
	}
}
