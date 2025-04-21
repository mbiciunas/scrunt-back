package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt"
	"scrunt-back/models/scrunt/script"
	"time"
)

type data struct {
	Name      string              `json:"name" binding:"required"`
	IconCode  string              `json:"icon_code" binding:"required"`
	DescShort string              `json:"desc_short" binding:"required"`
	DescLong  string              `json:"desc_long" binding:"required"`
	Source    scrunt.ScriptSource `json:"source" binding:"required"`
	Parent    string              `json:"parent"`
}

func PostScript(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		created := time.Now()
		uuid := script.GenerateScriptUUID(json.Name, created)

		id, err := script.GormInsertScript(uuid, json.Name, json.IconCode, json.DescShort, json.DescLong, json.Source, json.Parent, created)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		fmt.Println("api.scrunt.script.PostScript.go", "PostScript", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
