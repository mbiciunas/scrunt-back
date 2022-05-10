package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
)

type data struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Code        string `form:"code" json:"code" binding:"required"`
}

func PostScript(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Description: ", json.Description)
		fmt.Println("Script: ", json.Code)

		id, err := models.InsertScript(json.Name, json.Description, json.Code)
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
