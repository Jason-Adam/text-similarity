package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jason-adam/text-similarity/internal/api/controllers"
	"github.com/jason-adam/text-similarity/internal/api/middlewares"
)

// Setup sets up our middleware and routes
func Setup() *gin.Engine {
	app := gin.New()

	// Middlewares
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middlewares.CORS())
	app.NoRoute(middlewares.NoRouteHandler())

	// Routes
	app.GET("/api/v1/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	app.POST("/api/v1/similarity", controllers.CreateSimilarity)

	return app
}
