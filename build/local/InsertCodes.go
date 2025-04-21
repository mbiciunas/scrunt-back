package local

import (
	"fmt"
	"scrunt-back/models/scrunt/code"
)

var codeIdImport uint
var codeIdParameter uint
var codeIdMain uint

func InsertCodes() (err error) {
	fmt.Println("Insert Codes")

	codeIdImport = insertCode("import", "import libscrunt\\import libtest")
	codeIdParameter = insertCode("parameter", "ask('name', 'bla', 'bla')")
	codeIdMain = insertCode("main", "This is some code...")

	return err
}

func insertCode(codeType string, value string) uint {
	uuid := code.GenerateCodeUUID(codeType, value)
	codeId, err := code.GormInsertCode(uuid, codeType, value)

	if err != nil {
		panic(err)
	}

	return codeId
}
