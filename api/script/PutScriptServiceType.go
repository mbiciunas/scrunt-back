package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/scriptservicetype"
	"strconv"
)

type ScriptServiceTypeData struct {
	ScriptId      int    `json:"ScriptId" binding:"required"`
	ServiceTypeId int    `json:"ServiceTypeId" binding:"required"`
	Name          string `json:"Name" binding:"required"`
}

func PutScriptServiceType(c *gin.Context) {
	var json ScriptServiceTypeData

	fmt.Println("c.Keys: ", c.Keys)
	fmt.Println("c.Params: ", c.Params)
	scriptServiceTypeId, err := strconv.Atoi(c.Params.ByName("scriptservicetypeid"))
	if err != nil || scriptServiceTypeId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("scriptServiceTypeId: ", scriptServiceTypeId)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("ScriptId: ", json.ScriptId)
		fmt.Println("ServiceTypeId: ", json.ServiceTypeId)
		fmt.Println("Name: ", json.Name)

		rows, err := scriptservicetype.UpdateScriptServiceType(scriptServiceTypeId, json.ScriptId, json.ServiceTypeId, json.Name)
		if err != nil || rows != 1 {
			fmt.Println("PutScriptServiceType.PutScriptServiceType - err:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		fmt.Println("err.Error(): ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
