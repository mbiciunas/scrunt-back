package local

import (
	"fmt"
	"github.com/google/uuid"
	"scrunt-back/models/scrunt/code"
)

// ScriptTags
var codeIdImport uint
var codeIdParameter uint
var codeIdMain uint

func InsertCodes() (err error) {
	fmt.Println("Insert Codes")

	codeIdImport, err = code.GormInsertCode("import", "import libscrunt\\nimport libtest", genCodeUUID("import", "import libscrunt\\nimport libtest"))
	codeIdParameter, err = code.GormInsertCode("parameter", "ask('name', 'bla', 'bla')", genCodeUUID("parameter", "ask('name', 'bla', 'bla')"))
	codeIdMain, err = code.GormInsertCode("main", "This is some code...", genCodeUUID("main", "This is some code..."))

	return err
}

func genCodeUUID(codeType string, value string) string {
	fmt.Println("genCodeUUID", codeType, value)

	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(codeType+value)).String()
}
