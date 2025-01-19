package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/service"
	"strconv"
)

type Data struct {
	Name          string `json:"Name" binding:"required"`
	Description   string `json:"Description"`
	Address       string `json:"Address" binding:"required"`
	Port          int    `json:"Port" binding:"required"`
	ServiceTypeId int    `json:"ServiceTypeId" binding:"required"`
	CloudId       int    `json:"CloudId" binding:"-"`
}

func PutService(c *gin.Context) {
	var json Data

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("Id: ", id)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Description: ", json.Description)
		fmt.Println("Address: ", json.Address)
		fmt.Println("Port: ", json.Address)
		fmt.Println("ServiceTypeId: ", json.ServiceTypeId)
		fmt.Println("CloudId: ", json.CloudId)

		rows, err := service.UpdateService(id, json.Name, json.Description, json.Address, json.Port, json.ServiceTypeId, json.CloudId)
		if err != nil || rows != 1 {
			fmt.Println("PutService.PutService - err:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		fmt.Println("err.Error(): ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
