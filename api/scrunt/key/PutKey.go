package key

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/key"
	"strconv"
)

type Data struct {
	Name        string `json:"Name" binding:"required"`
	Description string `json:"Description"`
	Type        int64  `json:"Type" binding:"required"`
	KeyPublic   string `json:"KeyPublic" binding:"required"`
	KeyPrivate  string `json:"KeyPrivate" binding:"required"`
}

func PutKey(c *gin.Context) {
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
		fmt.Println("Type: ", json.Type)
		fmt.Println("KeyPublic: ", json.KeyPublic)
		fmt.Println("KeyPrivate: ", json.KeyPrivate)

		rows, err := key.UpdateKey(id, json.Name, json.Description, json.Type, json.KeyPublic, json.KeyPrivate)
		if err != nil || rows != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

	} else {
		fmt.Println("err.Error(): ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
