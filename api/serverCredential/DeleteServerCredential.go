package serverCredential

import (
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
	"strconv"
)

func DeleteServerCredential(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := models.DeleteServerCredential(id)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, rows)
	}
}
