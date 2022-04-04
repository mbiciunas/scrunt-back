package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
)

// Retrieve a list of all servers
func GetAllServer(c *gin.Context) {
	servers, err := models.SelectServersAll()
	fmt.Println("Server.GetAllServer - servers", servers)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, servers)
	}
}
