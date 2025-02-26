package script

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/runtime"
	"scrunt-back/models/store/script"
	"scrunt-back/python"
	"strconv"
)

type RunId struct {
	Id int64
}

func PostScriptRun(c *gin.Context) {
	var resultRunId RunId

	scriptId, err := strconv.Atoi(c.Param("scriptId"))
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	runId, err := runtime.InsertRuns(scriptId, 1)

	scriptDetail, err := script.SelectScript(scriptId)

	python.Exec(int(runId), scriptDetail.Code)

	resultRunId.Id = runId

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, resultRunId)
	}
}
