package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
)

// GetAllService Retrieve a list of all services
func GetAllService(c *gin.Context) {
	services, err := models.SelectServicesAll()
	fmt.Println("Service.GetAllService - services", services)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, services)
	}
}
