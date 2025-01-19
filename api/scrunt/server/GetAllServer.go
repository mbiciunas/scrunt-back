package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/server"
)

// GetAllServer Retrieve a list of all servers
func GetAllServer(c *gin.Context) {
	servers, err := server.SelectServersAll()
	fmt.Println("Server.GetAllServer - servers", servers)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, servers)
	}
}
