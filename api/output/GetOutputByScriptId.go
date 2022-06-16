package output

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
	"strconv"
)

func GetScript(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	output, err := models.SelectOutputByScriptId(id)
	jsonOutput, err := json.Marshal(output)
	fmt.Println(string(jsonOutput))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, output)
	}
}
