package script

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"scrunt-back/models/scrunt/script"
)

func GetScriptService(c *gin.Context) {
	scriptUUID := c.Param("scriptUUID")
	if err := uuid.Validate(scriptUUID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	scriptService, err := script.SelectScriptServices(scriptUUID)
	jsonScript, err := json.Marshal(scriptService)
	fmt.Println(string(jsonScript))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, scriptService)
	}
}
