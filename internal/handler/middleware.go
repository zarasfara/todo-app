package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonHeaderCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		acceptHeader := c.GetHeader("Accept")
		if acceptHeader != "application/json" {
			c.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported media type"})
			c.Abort()
		}
		c.Next()
	}
}
