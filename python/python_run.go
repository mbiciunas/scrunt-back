package python

import (
	"fmt"
	"github.com/go-python/cpy3"
	"strconv"
)

func Run(id int, code string) {
	//scriptStart := ""
	scriptStart := "from scrunt import Scrunt\nscrunt = Scrunt(" + strconv.Itoa(id) + ")\n\n"
	//scriptStart := "import sys\nsys.stdout = open('Output.txt', 'w')\n"
	//scriptEnd := "sys.stdout.close()"
	script := scriptStart + "\n" + code
	fmt.Println("RunId: ", id, "Script: ", script)
	//fmt.Println("Script: ", code)

	err := python3.PySys_SetArgv([]string{strconv.Itoa(id)})
	if err != nil {
		return
	}

	fmt.Println("Before Run")
	output := python3.PyRun_SimpleString(script)
	fmt.Println("After Run", output)
}
