package script

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
	"strconv"
)

func GetScriptService(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	scriptService, err := script.SelectScriptServices(id)
	jsonScript, err := json.Marshal(scriptService)
	fmt.Println(string(jsonScript))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, scriptService)
	}
}
