package middleware

import (
	"github.com/gin-gonic/gin"
)

// Logger 中间件举例
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
