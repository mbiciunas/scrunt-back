package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/servicescript"
	"strconv"
)

func GetServiceScript(c *gin.Context) {
	fmt.Println("GetServiceScript")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	serviceKey, err := servicescript.SelectServiceScripts(id)
	jsonScript, err := json.Marshal(serviceKey)
	fmt.Println(string(jsonScript))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, serviceKey)
	}
}
