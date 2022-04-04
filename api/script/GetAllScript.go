package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
)

func GetAllScript(c *gin.Context) {
	scripts, err := models.SelectScriptsAll()
	fmt.Println("Script.GetAllScript - scripts", scripts)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, scripts)
	}
}
