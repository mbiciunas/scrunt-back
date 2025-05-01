package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt"
	"scrunt-back/models/scrunt/script"
	"scrunt-back/models/scrunt/version"
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
	created := time.Now()
	major := uint(0)
	minor := uint(0)
	patch := uint(0)
	save := uint(0)
	change := "New Script"

	if err := c.ShouldBindJSON(&json); err == nil {
		uuidScript := script.GenerateScriptUUID(json.Name, created)

		scriptId, err := script.GormInsertScript(uuidScript, json.Name, json.IconCode, json.DescShort, json.DescLong, json.Source, json.Parent, created)
		if err != nil || scriptId <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		uuidVersion := version.GenerateVersionUUID(scriptId, major, minor, patch, save, change)

		id, err := version.GormInsertVersion(uuidVersion, scriptId, created, major, minor, patch, save, change)
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
