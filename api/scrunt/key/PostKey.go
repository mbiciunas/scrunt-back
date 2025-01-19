package key

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/key"
)

type data struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	KeyType     uint   `form:"keytype" json:"keytype" binding:"required"`
	KeyPrivate  string `form:"keyprivate" json:"keyprivate" binding:"required"`
	KeyPublic   string `form:"keypublic" json:"keypublic" binding:"required"`
}

func PostKey(c *gin.Context) {
	var json data
	fmt.Println("json: ", json)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Description: ", json.Description)
		fmt.Println("Type: ", json.KeyType)
		fmt.Println("Private: ", json.KeyPrivate)
		fmt.Println("Public: ", json.KeyPublic)

		id, err := key.InsertKey(json.Name, json.Description, json.KeyType, json.KeyPrivate, json.KeyPublic)
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
