package credential

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goginreact/models"
	"net/http"
)

// Retrieve a list of all credentials
func GetAllCredential(c *gin.Context) {
	credentials, err := models.SelectCredentialsAll()
	fmt.Println("Credential.GetAllCredential - credentials", credentials)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, credentials)
	}
}
