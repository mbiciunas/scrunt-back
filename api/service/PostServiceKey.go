package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/service"
)

type dataServiceKey struct {
	KeyId     uint `form:"keyid" json:"keyid" binding:"required"`
	ServiceId uint `form:"serviceid" json:"serviceid" binding:"required"`
}

func PostServiceKey(c *gin.Context) {
	var json dataServiceKey

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("PostServiceKey.PostServiceKey - json.KeyId:", json.KeyId, "json.ServiceId", json.ServiceId)
		fmt.Println("KeyId: ", json.KeyId)
		fmt.Println("ServiceId: ", json.ServiceId)

		id, err := service.InsertServiceKey(json.KeyId, json.ServiceId)
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
