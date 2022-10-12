package script

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/runtime"
	"scrunt-back/models/scrunt/script"
	"scrunt-back/python"
	"strconv"
)

//type data struct {
//	Name  string `form:"name" json:"name" binding:"required"`
//	Script  string `form:"script" json:"script" binding:"required"`
//}

type RunId struct {
	Id int64
}

func PostScriptRun(c *gin.Context) {
	//var json data
	var resultRunId RunId

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	run_id, err := runtime.InsertRuns(id, 1)
	//jsonScript, err := json.Marshal(script)

	script, err := script.SelectScript(id)
	//fmt.Println("PostScriptRun", "code: ", script.Code)
	//jsonScript, err := json.Marshal(script)

	//fmt.Println("PostScriptRun", "code: ", script.Code)
	//fmt.Println("PostScriptRun", "jsonScript: ", string(jsonScript))

	python.Run(int(run_id), script.Code)

	//if err := c.ShouldBindJSON(&json); err == nil {
	//	fmt.Println("Run Id: ", id)
	//	fmt.Println("Run Name: ", json.Name)
	//	fmt.Println("Run Code: ", json.Code)
	//
	//	//id, err := models.InsertScript(json.Name, json.Script)
	//	//if err != nil || id <= 0 {
	//	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	//	return
	//	//}
	//
	//	c.Header("Content-Type", "application/json")
	//	c.JSON(http.StatusOK, id)
	//
	//} else {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//}

	resultRunId.Id = run_id

	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, resultRunId)
	}
}

//func GetScript(c *gin.Context) {
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil || id < 1 {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err})
//		return
//	}
//
//	script, err := script.SelectScript(id)
//	jsonScript, err := json.Marshal(script)
//	fmt.Println(string(jsonScript))
//
//	if err == nil {
//		c.Header("Content-Type", "application/json")
//		c.JSON(http.StatusOK, script)
//	}
//}
