package script

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/store/rating"
	"scrunt-back/models/store/script"
	"scrunt-back/models/store/service"
	"scrunt-back/models/store/version"
	"strconv"
)

type payload struct {
	Script   script.GormScript
	Version  version.GormVersionAll
	Services []service.GormService
	Ratings  []rating.GormRating
}

func GetScript(c *gin.Context) {
	payload := payload{}

	scriptId, err := strconv.Atoi(c.Param("scriptId"))
	fmt.Println("store.script.GetScript - id = ", scriptId)
	if err != nil || scriptId < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	payload.Script, err = script.GormSelectScript(scriptId)
	if err != nil {
		fmt.Println("store.script.GetScript:  err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	payload.Version, err = version.GormSelectVersionNewest(scriptId)
	if err != nil {
		fmt.Println("store.script.GetScript:  err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println("store.script.GetScript:  payload.Version.Id = ", payload.Version.Id)

	payload.Services, err = service.GormSelectServicesForVersion(payload.Version.Id)
	if err != nil {
		fmt.Println("store.script.GetScript:  err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	payload.Ratings, err = rating.GormSelectRatingSummary(scriptId)
	if err != nil {
		fmt.Println("store.script.GetScript:  err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println("store.script.GetScript - payload: ", payload)

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("store.script.GetScript - err = ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("store.script.GetScript - jsonPayload: ", string(jsonPayload))

	//if err == nil {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, payload)
	//}
}
