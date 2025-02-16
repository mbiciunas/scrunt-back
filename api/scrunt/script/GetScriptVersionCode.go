package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/code"
	"strconv"
)

func GetScriptVersionCode(c *gin.Context) {
	scriptId, err := strconv.Atoi(c.Param("scriptId"))
	fmt.Println("store.script.GetScriptVersionCode - scriptId = ", scriptId)
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	versionId, err := strconv.Atoi(c.Param("versionId"))
	fmt.Println("store.script.GetScriptVersionCode - versionId = ", scriptId)
	if err != nil || versionId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	codes, err := code.GormSelectCodeForVersion(scriptId, versionId)
	fmt.Println("store.script.GetScriptVersionCode - codes", codes)
	if err != nil {
		fmt.Println("store.script.GetScriptVersionCode - err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, codes)
}
