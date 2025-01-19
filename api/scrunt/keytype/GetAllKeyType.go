package keytype

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/keytype"
)

// GetAllKeyType Retrieve a list of all key types
func GetAllKeyType(c *gin.Context) {
	keyTypes, err := keytype.SelectKeyTypesAll()
	fmt.Println("Keytype.GetAllKeyType - keyTypes", keyTypes)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, keyTypes)
	}
}
