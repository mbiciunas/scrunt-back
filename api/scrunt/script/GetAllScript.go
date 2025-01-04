package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
)

func GetAllScript(c *gin.Context) {
	scripts, err := script.SelectScriptsAll()
	fmt.Println("Script.GetAllScript - scripts", scripts)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, scripts)
	}
}
