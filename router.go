package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Allow access control
	router.Use(cors())

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile(".scrunt/frontend", true)))

	return router
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5174")
		//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
