package script

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/store/rating"
	"scrunt-back/models/store/script"
	"scrunt-back/models/store/service"
	"scrunt-back/models/store/version"
	"strconv"
)

type payload struct {
	Script   script.GormScript      `json:"script"`
	Version  version.GormVersionAll `json:"version"`
	Services *[]service.GormService `json:"service,omitempty"`
	Ratings  *[]rating.GormRating   `json:"rating,omitempty"`
}

func GetScript(c *gin.Context) {
	payload := payload{}

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

	payload.Version, err = version.GormSelectVersionNewest(scriptId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if payload.Version.Id > 0 {
		scriptServices, err := service.GormSelectServicesForVersion(payload.Version.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if len(scriptServices) > 0 {
			payload.Services = &scriptServices
		}
	}

	scriptRatings, err := rating.GormSelectRatingSummary(scriptId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if len(scriptRatings) > 0 {
		payload.Ratings = &scriptRatings
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, payload)
}
