package router

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jason-adam/text-similarity/internal/api/controllers"
	"github.com/jason-adam/text-similarity/internal/api/middlewares"
)

// Setup sets up our middleware and routes
func Setup() *gin.Engine {
	app := gin.New()

	// Logging to file
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Middlewares
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
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
