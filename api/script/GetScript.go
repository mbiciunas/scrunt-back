package script

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
	"strconv"
)

func GetScript(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println("id = " + string(id))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	scriptData, err := script.SelectScript(id)
	jsonScript, err := json.Marshal(scriptData)
	fmt.Println(string(jsonScript))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, scriptData)
	}
}
