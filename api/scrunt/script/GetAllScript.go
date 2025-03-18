package script

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
)

func GetAllScript(c *gin.Context) {
	scripts, err := script.GormSelectScriptsAll()
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, scripts)
	}
}
