package servicetype

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
)

// GetAllServiceType Retrieve a list of all service types
func GetAllServiceType(c *gin.Context) {
	serviceTypes, err := models.SelectServiceTypesAll()
	fmt.Println("Servicetype.GetAllServiceType - serviceTypes", serviceTypes)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, serviceTypes)
	}
}
