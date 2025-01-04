package project

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/models/scrunt/project"
)

// GetAllProject Retrieve a list of all projects
func GetAllProject(c *gin.Context) {
	projects, err := project.SelectProjectAll()
	fmt.Println("Project.GetAllProject - projects", projects)
	if err == nil {
		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, projects)
	}
}
