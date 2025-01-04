package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
	"strconv"
)

type Data struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description"`
	Code        string `json:"Code" binding:"required"`
}

func PutScript(c *gin.Context) {
	var json Data

	scriptId, err := strconv.Atoi(c.Params.ByName("scriptId"))
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("Id: ", scriptId)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Description: ", json.Description)
		fmt.Println("Code: ", json.Code)

		rows, err := script.UpdateScript(scriptId, json.Name, json.Description, json.Code)
		if err != nil || rows != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		fmt.Println("err.Error(): ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
