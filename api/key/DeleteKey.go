package key

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/key"
	"strconv"
)

func DeleteKey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := key.DeleteKey(id)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, rows)
	}
}
