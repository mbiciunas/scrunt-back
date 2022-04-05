package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
	"strconv"
)

// GetServer Retrieve a single server
func GetServer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	server, err := models.SelectServer(id)
	if err == nil && server.Id > 0 {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, server)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
}
