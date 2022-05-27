package output

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
)

type data struct {
	ScriptId int    `form:"script_id" json:"script_id" binding:"required"`
	Output   string `form:"output" json:"output" binding:"required"`
	Log      string `form:"log" json:"log" binding:"required"`
}

func PostOutput(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("ScriptId: ", json.ScriptId)
		fmt.Println("Output: ", json.Output)
		fmt.Println("Log: ", json.Log)

		id, err := models.InsertOutput(json.ScriptId, json.Output, json.Log)
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
