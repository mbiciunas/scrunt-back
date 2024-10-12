package key

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/servicekey"
	"strconv"
)

type ServiceKeyData struct {
	ServiceId int `json:"ServiceId" binding:"required"`
	KeyId     int `json:"KeyId" binding:"required"`
}

func PutKeyService(c *gin.Context) {
	var json ServiceKeyData

	fmt.Println("c.Keys: ", c.Keys)
	fmt.Println("c.Params: ", c.Params)
	serviceKeyId, err := strconv.Atoi(c.Params.ByName("servicekeyid"))
	if err != nil || serviceKeyId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("serviceKeyId: ", serviceKeyId)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("ServiceId: ", json.ServiceId)
		fmt.Println("KeyId: ", json.KeyId)

		rows, err := servicekey.UpdateServiceKey(serviceKeyId, json.ServiceId, json.KeyId)
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
