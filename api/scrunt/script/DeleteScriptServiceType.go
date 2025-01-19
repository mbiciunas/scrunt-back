package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/scriptservicetype"
	"strconv"
)

func DeleteScriptServiceType(c *gin.Context) {
	scriptServiceTypeId, err := strconv.Atoi(c.Param("scriptservicetypeid"))
	//serviceKeyId, err := strconv.Atoi(c.Param("servicekeyid"))
	//fmt.Println("serviceId", serviceId)
	fmt.Println("scriptServiceTypeId", scriptServiceTypeId)
	if err != nil || scriptServiceTypeId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	rows, err := scriptservicetype.DeleteScriptServiceType(scriptServiceTypeId)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, rows)
	}
}
