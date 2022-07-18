package python

import (
	"github.com/go-python/cpy3"
)

func Finalize() {
	python3.Py_Finalize()

	python3.PyErr_Print()
}
