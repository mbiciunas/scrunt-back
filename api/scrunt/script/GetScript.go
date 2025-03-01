package script

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/script"
	"scrunt-back/models/scrunt/service"
	"scrunt-back/models/scrunt/version"
	"strconv"
)

type Payload struct {
	Script   script.GormScript       `json:"script"`
	Version  *version.GormVersionAll `json:"version,omitempty"`
	Services *[]service.GormService  `json:"service,omitempty"`
}

func GetScript(c *gin.Context) {
	payload := Payload{}

	scriptId, err := strconv.Atoi(c.Param("scriptId"))
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	payload.Script, err = script.GormSelectScript(scriptId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	scriptVersion, err := version.GormSelectVersionNewest(scriptId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if scriptVersion.Id > 0 {
		payload.Version = &scriptVersion

		scriptServices, err := service.GormSelectServicesForVersion(payload.Version.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if len(scriptServices) > 0 {
			payload.Services = &scriptServices
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, payload)
}
