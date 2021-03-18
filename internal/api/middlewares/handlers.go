package middlewares

import "github.com/gin-gonic/gin"

// NoMethodHandler returns an error message if the method is not permitted
func NoMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "method not allowed"})
	}
}

// NoRouteHandler returns an error message if function of route was not found
func NoRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "The processing function of the request route was not found"})
	}
}
