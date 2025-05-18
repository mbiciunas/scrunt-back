package main

import (
	"github.com/gin-gonic/gin"
	"scrunt-back/api/scrunt/cloud"
	"scrunt-back/api/scrunt/credential"
	"scrunt-back/api/scrunt/icon"
	"scrunt-back/api/scrunt/key"
	"scrunt-back/api/scrunt/keytype"
	"scrunt-back/api/scrunt/output"
	"scrunt-back/api/scrunt/project"
	"scrunt-back/api/scrunt/script"
	"scrunt-back/api/scrunt/server"
	"scrunt-back/api/scrunt/serverCredential"
	"scrunt-back/api/scrunt/service"
	"scrunt-back/api/scrunt/servicetype"
)

func apiScrunt(router *gin.Engine) {
	scruntClouds(router)
	scruntCredentials(router)
	scruntIcons(router)
	scruntKeys(router)
	scruntKeyTypes(router)
	scruntProjects(router)
	scruntScripts(router)
	scruntServices(router)
	scruntServiceTypes(router)
	scruntServers(router)
	scruntServerCredentials(router)
	scruntOutputs(router)
}

func scruntClouds(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/clouds", cloud.GetAllCloud)
	}
}

func scruntProjects(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/projects", project.GetAllProject)
	}
}

func scruntCredentials(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/credentials", credential.GetAllCredential)
		api.POST("/credentials", credential.PostCredential)
		api.GET("/credentials/:id", credential.GetCredential)
		api.PUT("/credentials/:id", credential.PutCredential)
		api.DELETE("/credentials/:id", credential.DeleteCredential)
	}
}

func scruntIcons(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/icons", icon.GetAllIcon)
	}
}

func scruntKeys(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/keys", key.GetAllKey)
		api.POST("/keys", key.PostKey)
		api.GET("/keys/:id", key.GetKey)
		api.PUT("/keys/:id", key.PutKey)
		api.GET("/keys/:id/services", key.GetKeyService)
		api.POST("/keys/:id/services", key.PostKeyService)
		api.DELETE("/keys/:id", key.DeleteKey)
		api.PUT("/keys/:id/services/:servicekeyid", key.PutKeyService)
	}
}

func scruntKeyTypes(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/keytypes", keytype.GetAllKeyType)
	}
}

func scruntServices(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/services", service.GetAllService)
		api.POST("/services", service.PostService)
		api.GET("/services/:id", service.GetService)
		api.PUT("/services/:id", service.PutService)
		api.GET("/services/:id/keys", service.GetServiceKey)
		api.POST("/services/:id/keys", service.PostServiceKey)
		api.DELETE("/services/:id/keys/:servicekeyid", service.DeleteServiceKey)
		api.PUT("/services/:id/keys/:servicekeyid", service.PutServiceKey)
		api.GET("/services/:id/scripts", service.GetServiceScript)
	}
}

func scruntServiceTypes(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/servicetypes", servicetype.GetAllServiceType)
	}
}

func scruntServers(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/servers", server.GetAllServer)
		api.POST("/servers", server.PostServer)
		api.GET("/servers/:id", server.GetServer)
		api.PUT("/servers/:id", server.PutServer)
		api.DELETE("/servers/:id", server.DeleteServer)
	}
}

func scruntServerCredentials(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.DELETE("/servercredentials/:id", serverCredential.DeleteServerCredential)
	}
}

func scruntScripts(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.GET("/scripts", script.GetAllScript)
		api.POST("/scripts", script.PostScript)
		api.GET("/scripts/:scriptUUID", script.GetScript)
		api.PUT("/scripts/:scriptId", script.PutScript)
		api.DELETE("/scripts/:scriptId", script.DeleteScript)
		api.POST("/scripts/:scriptId/run", script.PostScriptRun)
		api.GET("/scripts/:scriptUUID/issues", script.GetScriptIssueAll)
		api.GET("/scripts/:scriptUUID/services", script.GetScriptService)
		api.POST("/scripts/:scriptId/services", script.PostScriptServiceType)
		api.DELETE("/scripts/:scriptId/services/:scriptservicetypeid", script.DeleteScriptServiceType)
		api.PUT("/scripts/:scriptId/services/:scriptservicetypeid", script.PutScriptServiceType)
		api.GET("/scripts/:scriptUUID/versions", script.GetScriptVersionAll)
		api.GET("/scripts/:scriptUUID/versions/:versionId/codes", script.GetScriptVersionCode)
	}
}

func scruntOutputs(router *gin.Engine) {
	api := router.Group("/api/scrunt")
	{
		api.POST("/outputs", output.PostOutput)
		api.GET("/outputs/script/:runid/:id", output.GetOutputByRunId)
	}
}
