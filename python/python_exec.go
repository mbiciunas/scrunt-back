package python

import (
	"fmt"
	"os"

	//"github.com/go-python/cpy3"
	"log"
	"os/exec"
	"strconv"
)

func Exec(id int, code string) {
	//scriptStart := ""
	scriptStart := "from libscrunt import Scrunt\n"
	scriptStart += "scrunt = Scrunt(" + strconv.Itoa(id) + ")\n\n"
	//scriptStart := "import sys\nsys.stdout = open('Output.txt', 'w')\n"
	//scriptEnd := "sys.stdout.close()"
	script := scriptStart + "\n" + code
	fmt.Println("RunId: ", id, "Script: ", script)
	//fmt.Println("Script: ", code)

	//err := python3.PySys_SetArgv([]string{strconv.Itoa(id)})
	//if err != nil {
	//	return
	//}

	//
	// Create temp file in local directory
	//
	f, err := os.CreateTemp(".", "*.py")
	if err != nil {
		log.Fatalf("Failed to create temp file", err)
	}
	fmt.Println("Temp file name:", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte(script))
	if err != nil {
		log.Fatalf("Failed to write to temp file", err)
	}

	fmt.Println("Before Exec")
	//output := python3.PyRun_SimpleString(script)
	//cmd := exec.Command("ls", "-lah")
	//cmd := exec.Command(".scrunt/python3/bin/python3", "-I", f.Name())
	cmd := exec.Command(".scrunt/python/bin/python3", f.Name())
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n\n%s", err, out)
	}
	//fmt.Printf("combined out:\n%s\n", string(out))
	//fmt.Println("After Exec", output)
}

//python -I -c "commands"
