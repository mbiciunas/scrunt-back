package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"scrunt-back/api/cloud"
	"scrunt-back/api/credential"
	"scrunt-back/api/key"
	"scrunt-back/api/keytype"
	"scrunt-back/api/output"
	"scrunt-back/api/project"
	"scrunt-back/api/script"
	"scrunt-back/api/server"
	"scrunt-back/api/serverCredential"
	"scrunt-back/api/service"
	"scrunt-back/api/servicetype"
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

	api.GET("/clouds", cloud.GetAllCloud)

	api.GET("/projects", project.GetAllProject)

	api.GET("/credentials", credential.GetAllCredential)
	api.POST("/credentials", credential.PostCredential)
	api.GET("/credentials/:id", credential.GetCredential)
	api.PUT("/credentials/:id", credential.PutCredential)
	api.DELETE("/credentials/:id", credential.DeleteCredential)

	api.GET("/keys", key.GetAllKey)
	api.POST("/keys", key.PostKey)
	api.GET("/keys/:id", key.GetKey)
	api.PUT("/keys/:id", key.PutKey)
	api.GET("/keys/:id/services", key.GetKeyService)
	api.POST("/keys/:id/services", key.PostKeyService)
	api.DELETE("/keys/:id", key.DeleteKey)
	api.PUT("/keys/:id/services/:servicekeyid", key.PutKeyService)

	api.GET("/keytypes", keytype.GetAllKeyType)

	api.GET("/services", service.GetAllService)
	api.POST("/services", service.PostService)
	api.GET("/services/:id", service.GetService)
	api.PUT("/services/:id", service.PutService)
	api.GET("/services/:id/keys", service.GetServiceKey)
	api.POST("/services/:id/keys", service.PostServiceKey)
	api.DELETE("/services/:id/keys/:servicekeyid", service.DeleteServiceKey)
	api.PUT("/services/:id/keys/:servicekeyid", service.PutServiceKey)
	api.GET("/services/:id/scripts", service.GetServiceScript)

	api.GET("/servicetypes", servicetype.GetAllServiceType)

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
	api.DELETE("/scripts/:id", script.DeleteScript)
	api.POST("/scripts/:id/run", script.PostScriptRun)
	api.GET("/scripts/:id/services", script.GetScriptService)
	api.POST("/scripts/:id/services", script.PostScriptServiceType)
	api.DELETE("/scripts/:id/services/:scriptservicetypeid", script.DeleteScriptServiceType)
	api.PUT("/scripts/:id/services/:scriptservicetypeid", script.PutScriptServiceType)

	//api.GET("/outputs/:id", script.GetOutput)
	api.POST("/outputs", output.PostOutput)
	api.GET("/outputs/script/:runid/:id", output.GetOutputByRunId)
}
