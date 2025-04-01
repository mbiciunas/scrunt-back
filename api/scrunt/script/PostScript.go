package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
	"time"
)

type data struct {
	Name      string `form:"name" json:"name" binding:"required"`
	IconCode  string `form:"icon" json:"icon_code" binding:"required"`
	DescShort string `form:"desc_short" json:"desc_short" binding:"required"`
	DescLong  string `form:"desc_long" json:"desc_long" binding:"required"`
}

func PostScript(c *gin.Context) {
	var json data
	//fmt.Println(c.)
	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Icon: ", json.IconCode)
		fmt.Println("Desc Short: ", json.DescShort)
		fmt.Println("Desc Long: ", json.DescLong)

		id, err := script.GormInsertScript(json.Name, json.IconCode, json.DescShort, json.DescLong, time.Now())
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
