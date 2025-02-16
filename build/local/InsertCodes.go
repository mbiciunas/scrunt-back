package local

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ScriptTags
var codeImport Code
var codeParameter Code
var codeMain Code

func InsertCodes(db *gorm.DB) {
	fmt.Println("Insert Codes")

	codeImport = Code{Type: "import", Value: "import libscrunt\\nimport libtest"}
	codeImport.Uuid = genCodeUUID(codeImport.Type, codeImport.Value)
	codeParameter = Code{Type: "parameter", Value: "ask('name', 'bla', 'bla')"}
	codeParameter.Uuid = genCodeUUID(codeParameter.Type, codeParameter.Value)
	codeMain = Code{Type: "main", Value: "This is some code..."}
	codeMain.Uuid = genCodeUUID(codeMain.Type, codeMain.Value)

	db.Create(&codeImport)
	db.Create(&codeParameter)
	db.Create(&codeMain)
}

func genCodeUUID(codeType string, value string) string {
	fmt.Println("genCodeUUID", codeType, value)

	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(codeType+value)).String()
}
