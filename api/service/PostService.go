package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/service"
)

type data struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Address     string `form:"address" json:"address" binding:"required"`
	Port        uint   `form:"port" json:"port" binding:"required"`
	ServiceType uint   `form:"servicetype" json:"servicetype" binding:"required"`
}

func PostService(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Description: ", json.Description)
		fmt.Println("Address: ", json.Address)
		fmt.Println("Port: ", json.Port)
		fmt.Println("ServiceType: ", json.ServiceType)

		//
		// The foreign relationship in services for cloud_id is pointing the wrong way (I think) and should be nullable.
		// Update the data model, reload the data (from scripts) and try adding a service.  Should now work.
		//
		id, err := service.InsertService(json.Name, json.Description, json.Address, json.Port, json.ServiceType)
		if err != nil || id <= 0 {
			fmt.Println("JSON: ", json)
			fmt.Println("Error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		fmt.Println("JSON: ", json)
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
