package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"scrunt-back/models/scrunt/version"
)

func GetScriptVersionAll(c *gin.Context) {
	scriptUUID := c.Param("scriptUUID")
	if err := uuid.Validate(scriptUUID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	versions, err := version.GormSelectVersionsAll(scriptUUID)
	fmt.Println("scrunt.script.GetScriptVersionAll - codes", versions)
	if err != nil {
		fmt.Println("scrunt.script.GetScriptVersionAll - err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, versions)
}
