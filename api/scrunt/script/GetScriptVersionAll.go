package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/version"
	"strconv"
)

func GetScriptVersionAll(c *gin.Context) {
	scriptId, err := strconv.Atoi(c.Param("scriptId"))
	fmt.Println("store.script.GetScriptVersionAll - scriptId = ", scriptId)
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	versions, err := version.GormSelectVersionsAll(scriptId)
	fmt.Println("store.script.GetScriptVersionAll - codes", versions)
	if err != nil {
		fmt.Println("store.script.GetScriptVersionAll - err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, versions)
}
