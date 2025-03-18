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
	DescShort string `form:"desc_short" json:"desc_short" binding:"required"`
	DescLong  string `form:"desc_long" json:"desc_long" binding:"required"`
}

func PostScript(c *gin.Context) {
	var json data
	//fmt.Println(c.)
	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Desc Short: ", json.DescShort)
		fmt.Println("Desc Long: ", json.DescLong)

		id, err := script.GormInsertScript(1, json.Name, json.DescShort, json.DescLong, time.Now())
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
