package credential

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
	"strconv"
)

// GetCredential Retrieve a single credential
func GetCredential(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	credential, err := models.SelectCredential(id)
	jsonCredential, err := json.Marshal(credential)
	fmt.Println(string(jsonCredential))

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, credential)
	}
}
