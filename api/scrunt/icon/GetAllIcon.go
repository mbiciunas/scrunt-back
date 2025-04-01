package icon

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/icon"
)

// GetAllIcon Retrieve a list of all clouds
func GetAllIcon(c *gin.Context) {
	icons, err := icon.GormSelectIconsAll()
	fmt.Println("Icon.GetAllIcon - icons", icons)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, icons)
	}
}
