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
		api.GET("/scripts/:scriptId", script.GetScript)
		api.PUT("/scripts/:scriptId", script.PutScript)
		api.DELETE("/scripts/:scriptId", script.DeleteScript)
		api.POST("/scripts/:scriptId/run", script.PostScriptRun)
		api.GET("/scripts/:scriptId/issues", script.GetScriptIssueAll)
		api.GET("/scripts/:scriptId/services", script.GetScriptService)
		api.POST("/scripts/:scriptId/services", script.PostScriptServiceType)
		api.GET("/scripts/:scriptId/versions", script.GetScriptVersionAll)
		api.GET("/scripts/:scriptId/versions/:versionId/codes", script.GetScriptVersionCode)
		api.DELETE("/scripts/:scriptId/services/:scriptservicetypeid", script.DeleteScriptServiceType)
		api.PUT("/scripts/:scriptId/services/:scriptservicetypeid", script.PutScriptServiceType)
	}
}
