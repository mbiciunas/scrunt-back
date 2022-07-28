package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/api/credential"
	"scrunt-back/api/output"
	"scrunt-back/api/script"
	"scrunt-back/api/server"
	"scrunt-back/api/serverCredential"
)

func api(router *gin.Engine) {
	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	api.GET("/credentials", credential.GetAllCredential)
	api.POST("/credentials", credential.PostCredential)
	api.GET("/credentials/:id", credential.GetCredential)
	api.PUT("/credentials/:id", credential.PutCredential)
	api.DELETE("/credentials/:id", credential.DeleteCredential)

	api.GET("/servers", server.GetAllServer)
	api.POST("/servers", server.PostServer)
	api.GET("/servers/:id", server.GetServer)
	api.PUT("/servers/:id", server.PutServer)
	api.DELETE("/servers/:id", server.DeleteServer)

	api.DELETE("/servercredentials/:id", serverCredential.DeleteServerCredential)

	api.GET("/scripts", script.GetAllScript)
	api.POST("/scripts", script.PostScript)
	api.GET("/scripts/:id", script.GetScript)
	api.PUT("/scripts/:id", script.PutScript)
	api.POST("/scripts/:id/run", script.PostScriptRun)
	api.DELETE("/scripts/:id", script.DeleteScript)

	//api.GET("/outputs/:id", script.GetOutput)
	api.POST("/outputs", output.PostOutput)
	api.GET("/outputs/script/:runid/:id", output.GetOutputByRunId)
}
