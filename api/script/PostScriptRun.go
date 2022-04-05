package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//type data struct {
//	Name  string `form:"name" json:"name" binding:"required"`
//	Script  string `form:"script" json:"script" binding:"required"`
//}

func PostScriptRun(c *gin.Context) {
	var json data

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Run Id: ", id)
		fmt.Println("Run Name: ", json.Name)
		fmt.Println("Run Script: ", json.Script)

		//id, err := models.InsertScript(json.Name, json.Script)
		//if err != nil || id <= 0 {
		//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
		//	return
		//}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
