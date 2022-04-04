package credential

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
)

type data struct {
	Name       string `form:"name" json:"name" binding:"required"`
	CredType   string `form:"credtype" json:"credtype" binding:"required"`
	Credential string `form:"credential" json:"credential" binding:"required"`
}

func PostCredential(c *gin.Context) {
	var json data

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Type: ", json.CredType)
		fmt.Println("credential: ", json.Credential)

		id, err := models.InsertCredential(json.Name, json.CredType, json.Credential)
		if err != nil || id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, id)

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
