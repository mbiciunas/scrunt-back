package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/service"
	"strconv"
)

func DeleteServiceKey(c *gin.Context) {
	serviceId, err := strconv.Atoi(c.Param("id"))
	serviceKeyId, err := strconv.Atoi(c.Param("servicekeyid"))
	fmt.Println("serviceId", serviceId)
	fmt.Println("serviceKeyId", serviceKeyId)
	if err != nil || serviceId < 1 || serviceKeyId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := service.DeleteServiceKey(serviceKeyId)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, rows)
	}
}
