package script

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models"
	"strconv"
)

type Data struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description"`
	Code        string `json:"Code" binding:"required"`
}

func PutScript(c *gin.Context) {
	var json Data

	//jsonData, err := ioutil.ReadAll(c.Request.Body)
	//if err != nil {
	//	// Handle error
	//}
	//fmt.Println("jsonData: ", string(jsonData))
	//
	////id := c.Params.ByName("id")
	//fmt.Println("ByName(\"Name\"): ", c.Params.ByName("Name"))
	//fmt.Println("ByName(\"Description\"): ", c.Params.ByName("Description"))
	//fmt.Println("ByName(\"Code\"): ", c.Params.ByName("Code"))
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	fmt.Println("Id: ", id)

	if err := c.ShouldBindJSON(&json); err == nil {
		fmt.Println("Name: ", json.Name)
		fmt.Println("Description: ", json.Description)
		fmt.Println("Code: ", json.Code)

		rows, err := models.UpdateScript(id, json.Name, json.Description, json.Code)
		if err != nil || rows != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		fmt.Println("err.Error(): ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
