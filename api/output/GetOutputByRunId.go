package output

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/runtime"
	"strconv"
)

func GetOutputByRunId(c *gin.Context) {
	fmt.Println("Enter GetOutputByRunId")
	runId, err := strconv.Atoi(c.Param("runid"))
	fmt.Println("runId", runId)
	if err != nil || runId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("id", id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	output, err := runtime.SelectOutputByRunId(runId, id)
	jsonOutput, err := json.Marshal(output)
	fmt.Println(string(jsonOutput))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, output)
	}
}
