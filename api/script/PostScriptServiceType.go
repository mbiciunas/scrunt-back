package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/scriptservicetype"
)

type scriptServiceTypeData struct {
	ScriptId      uint   `json:"ScriptId" binding:"required"`
	ServiceTypeId uint   `json:"ServiceTypeId" binding:"required"`
	Name          string `json:"Name" binding:"required"`
}

//type dataServiceKey struct {
//	KeyId     uint `form:"keyid" json:"keyid" binding:"required"`
//	ServiceId uint `form:"serviceid" json:"serviceid" binding:"required"`
//}

func PostScriptServiceType(c *gin.Context) {
	var json scriptServiceTypeData

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("PostScriptServiceType.PostServiceKey - json.ScriptId:", json.ScriptId, "json.ServiceTypeId", json.ServiceTypeId, "json.Name", json.Name)

		id, err := scriptservicetype.InsertScriptServiceType(json.ScriptId, json.ServiceTypeId, json.Name)
		if err != nil || id <= 0 {
			fmt.Println("JSON: ", json)
			fmt.Println("Error: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		fmt.Println("JSON: ", json)
		fmt.Println("Error: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
