package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"scrunt-back/models/scrunt/script"
	"scrunt-back/models/store/issue"
)

func GetScriptIssueAll(c *gin.Context) {
	var err error

	scriptUUID := c.Param("scriptUUID")
	if err := uuid.Validate(scriptUUID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	scriptDetails, err := script.GormSelectScript(scriptUUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	issues, err := issue.GormSelectIssues(scriptDetails.Parent)

	if err != nil {
		fmt.Println("store.script.GetScriptIssueAll - err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, issues)
}
