package middleware

import (
	"github.com/gin-gonic/gin"
)

// 中间件举例
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
