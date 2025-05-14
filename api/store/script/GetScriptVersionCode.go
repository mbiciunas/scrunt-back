package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"scrunt-back/models/store/code"
	"strconv"
)

func GetScriptVersionCode(c *gin.Context) {
	var err error

	scriptUUID := c.Param("scriptUUID")
	if err := uuid.Validate(scriptUUID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	versionId, err := strconv.Atoi(c.Param("versionId"))
	fmt.Println("store.script.GetScriptVersionCode - versionId = ", versionId)
	if err != nil || versionId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	codes, err := code.GormSelectCodeForVersion(scriptUUID, versionId)
	fmt.Println("store.script.GetScriptVersionCode - codes", codes)
	if err != nil {
		fmt.Println("store.script.GetScriptVersionCode - err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, codes)
}
