package middlewares

import (
	"io"
	"net/http"

	"github.com/fazanurfaizi/go-rest-template/pkg/sanitize"
	"github.com/gin-gonic/gin"
)

type SanitizeMiddleware struct{}

func NewSanitizeMiddleware() *SanitizeMiddleware {
	return &SanitizeMiddleware{}
}

func (m SanitizeMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
		}
		defer ctx.Request.Body.Close()

		sanitizeBody, err := sanitize.SanitizeJSON(body)
		if err != nil {
			ctx.Writer.WriteHeader(http.StatusBadRequest)
		}

		ctx.Set("body", sanitizeBody)
		ctx.Next()
	}
}
