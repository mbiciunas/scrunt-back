package credential

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/credential"
)

// GetAllCredential Retrieve a list of all credentials
func GetAllCredential(c *gin.Context) {
	credentials, err := credential.SelectCredentialsAll()
	fmt.Println("Credential.GetAllCredential - credentials", credentials)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, credentials)
	}
}
