package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/server"
	"strconv"
)

type Data struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
}

func PutServer(c *gin.Context) {
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
		fmt.Println("Address: ", json.Address)

		rows, err := server.UpdateServer(id, json.Name, json.Address)
		if err != nil || rows != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
