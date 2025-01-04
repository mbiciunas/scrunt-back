package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/store/script"
)

func GetAllScript(c *gin.Context) {
	scripts, err := script.GormSelectScriptsAll()
	fmt.Println("Script.GetAllScript - scripts", scripts)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, scripts)
	}
}
