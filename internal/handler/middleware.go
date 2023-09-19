package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func jsonHeaderCheckMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		acceptHeader := ctx.GetHeader("Accept")
		if acceptHeader != "application/json" {
			ctx.JSON(http.StatusUnsupportedMediaType, gin.H{"error": "Unsupported media type"})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func setJSONContentType() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		ctx.Next()
	}
}
