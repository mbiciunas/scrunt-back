package script

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
	"strconv"
)

func GetScript(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	script, err := models.SelectScript(id)
	jsonScript, err := json.Marshal(script)
	fmt.Println(string(jsonScript))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, script)
	}
}
