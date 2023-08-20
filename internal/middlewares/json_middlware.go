package middlewares

import "github.com/gin-gonic/gin"

type JsonMiddleware struct{}

func NewJsonMiddleware() JsonMiddleware {
	return JsonMiddleware{}
}

func (p JsonMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Next()
	}
}
