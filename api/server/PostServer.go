package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
)

type data struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
}

func PostServer(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Address: ", json.Address)

		id, err := models.InsertServer(json.Name, json.Address)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
