package main

import (
	"github.com/gin-gonic/gin"
	"scrunt-back/api/store/script"
)

func apiStore(router *gin.Engine) {
	storeScripts(router)
}

func storeScripts(router *gin.Engine) {
	api := router.Group("/api/store")
	{
		api.GET("/scripts", script.GetAllScript)
		api.POST("/scripts", script.PostScript)
		api.GET("/scripts/:scriptUUID", script.GetScript)
		api.PUT("/scripts/:scriptId", script.PutScript)
		api.DELETE("/scripts/:scriptId", script.DeleteScript)
		api.POST("/scripts/:scriptId/run", script.PostScriptRun)
		api.GET("/scripts/:scriptUUID/issues", script.GetScriptIssueAll)
		//api.GET("/scripts/:scriptUUID/services", script.GetScriptService)
		//api.POST("/scripts/:scriptUUID/services", script.PostScriptServiceType)
		api.DELETE("/scripts/:scriptId/services/:scriptservicetypeid", script.DeleteScriptServiceType)
		api.PUT("/scripts/:scriptId/services/:scriptservicetypeid", script.PutScriptServiceType)
		api.GET("/scripts/:scriptUUID/versions", script.GetScriptVersionAll)
		api.GET("/scripts/:scriptUUID/versions/:versionId/codes", script.GetScriptVersionCode)
	}
}
