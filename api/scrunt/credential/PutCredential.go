package credential

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/credential"
	"strconv"
)

type Data struct {
	Name       string `form:"name" json:"name" binding:"required"`
	Credential string `form:"credential" json:"credential" binding:"required"`
}

// PutCredential Update a single credential
func PutCredential(c *gin.Context) {
	var json Data

	//id := c.Params.ByName("id")
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("Id: ", id)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("credential: ", json.Credential)

		rows, err := credential.UpdateCredential(id, json.Name, json.Credential)
		if err != nil || rows != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
