package script

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/runtime"
	"scrunt-back/models/scrunt/script"
	"scrunt-back/python"
	"strconv"
)

type RunId struct {
	Id int64
}

func PostScriptRun(c *gin.Context) {
	var resultRunId RunId

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	runId, err := runtime.InsertRuns(id, 1)

	scriptDetail, err := script.SelectScript(id)

	python.Exec(int(runId), scriptDetail.Code)

	resultRunId.Id = runId

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, resultRunId)
	}
}
