package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/store/issue"
	"strconv"
)

func GetScriptIssueAll(c *gin.Context) {
	scriptId, err := strconv.Atoi(c.Param("scriptId"))
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	issues, err := issue.GormSelectIssues(scriptId)
	fmt.Println("store.issue.GetScriptIssueAll - issues", issues)
	if err != nil {
		fmt.Println("store.script.GetScriptIssueAll - err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, issues)
}
