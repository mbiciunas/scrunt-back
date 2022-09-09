package cloud

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/cloud"
)

// GetAllCloud Retrieve a list of all clouds
func GetAllCloud(c *gin.Context) {
	clouds, err := cloud.SelectCloudAll()
	fmt.Println("Cloud.GetAllCloud - clouds", clouds)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, clouds)
	}
}
