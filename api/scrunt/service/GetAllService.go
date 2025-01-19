package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/service"
)

// GetAllService Retrieve a list of all services
func GetAllService(c *gin.Context) {
	services, err := service.SelectServicesAll()
	fmt.Println("service.GetAllService - services", services)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, services)
	}
}
