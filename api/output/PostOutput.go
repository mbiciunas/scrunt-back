package output

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/runtime"
)

type data struct {
	RunId      int    `form:"run_id" json:"run_id" binding:"required"`
	OutputType int    `form:"output_type" json:"output_type" binding:"required"`
	Value      string `form:"value" json:"value" binding:"required"`
}

func PostOutput(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		//fmt.Println("ScriptId: ", json.ScriptId)
		//fmt.Println("Output: ", json.Output)
		//fmt.Println("Log: ", json.Log)

		id, err := runtime.InsertOutput(json.RunId, json.OutputType, json.Value)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
