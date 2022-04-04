package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
	"strconv"
)

type Data struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Script string `form:"script" json:"script" binding:"required"`
}

func PutScript(c *gin.Context) {
	var json Data

	//id := c.Params.ByName("id")
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("Id: ", id)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Script: ", json.Script)

		rows, err := models.UpdateScript(id, json.Name, json.Script)
		if err != nil || rows != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
