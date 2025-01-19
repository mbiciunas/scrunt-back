package key

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/key"
	"strconv"
)

func GetKey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	keyData, err := key.SelectKey(id)
	jsonScript, err := json.Marshal(keyData)
	fmt.Println(string(jsonScript))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, keyData)
	}
}
