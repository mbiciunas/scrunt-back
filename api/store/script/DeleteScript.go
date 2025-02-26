package script

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/store/script"
	"strconv"
)

func DeleteScript(c *gin.Context) {
	scriptId, err := strconv.Atoi(c.Param("scriptId"))
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := script.DeleteScript(scriptId)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, rows)
	}
}
