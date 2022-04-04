package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
)

type data struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Script string `form:"script" json:"script" binding:"required"`
}

func PostScript(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Script: ", json.Script)

		id, err := models.InsertScript(json.Name, json.Script)
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
