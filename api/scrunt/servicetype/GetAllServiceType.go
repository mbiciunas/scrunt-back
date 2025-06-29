package servicetype

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/servicetype"
)

func GetAllServiceType(c *gin.Context) {
	serviceTypes, err := servicetype.GormSelectServiceTypesAll()
	fmt.Println("Servicetype.GormSelectServiceTypesAll - serviceTypes", serviceTypes)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, serviceTypes)
	}
}
