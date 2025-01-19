package key

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/key"
)

// GetAllKey Retrieve a list of all keys
func GetAllKey(c *gin.Context) {
	keys, err := key.SelectKeysAll()
	fmt.Println("key.GetAllKey - keys", keys)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, keys)
	}
}
