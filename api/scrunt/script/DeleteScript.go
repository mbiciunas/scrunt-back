package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt"
	"scrunt-back/models/scrunt/script"
	"strconv"
)

func DeleteScript(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("scriptId"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tx := scrunt.GormDB.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		fmt.Println("api.scrunt.script.DeleteScript", "transaction error", err, tx.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := script.GormDeleteScript(tx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tx.Commit()
	if tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": tx.Error})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, rows)
}
