package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
	"strconv"
)

type Data struct {
	Id        uint   `form:"id" json:"id" binding:"required"`
	Name      string `form:"name" json:"name" binding:"required"`
	IconCode  string `form:"icon" json:"icon_code" binding:"required"`
	DescShort string `form:"desc_short" json:"desc_short" binding:"required"`
	DescLong  string `form:"desc_long" json:"desc_long" binding:"required"`
}

func PutScript(c *gin.Context) {
	var json Data

	fmt.Println("api.scrunt.script.PutScript", "Start: ")
	fmt.Println("api.scrunt.script.PutScript", "Id: ", c.Params.ByName("scriptId"))

	id, err := strconv.Atoi(c.Params.ByName("scriptId"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("Id: ", id)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Id: ", json.Id)
		fmt.Println("Name: ", json.Name)
		fmt.Println("IconCode: ", json.IconCode)
		fmt.Println("DescShort: ", json.DescShort)
		fmt.Println("DescLong: ", json.DescLong)

		rows, err := script.GormUpdateScript(json.Id, json.Name, json.IconCode, json.DescShort, json.DescLong)
		if err != nil || rows != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		fmt.Println("err.Error(): ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
