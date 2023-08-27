package middlewares

import (
	"net/http"

	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/csrf"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CsrfMiddleware struct {
	config *config.Config
	logger logger.Logger
}

func NewCsrfMiddleware(config *config.Config, logger logger.Logger) *CsrfMiddleware {
	return &CsrfMiddleware{config: config, logger: logger}
}

func (m CsrfMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !m.config.Server.CSRF {
			ctx.Next()
		}

		token := ctx.Request.Header.Get(csrf.CSRFHeader)
		if token == "" {
			m.logger.Errorf("CSRF Middleware get CSRF Header, Token: %s, Error: %s",
				token,
				"empty CSRF token",
			)
			utils.ErrorJSON(ctx, http.StatusForbidden, errors.NewRestError(http.StatusForbidden, "Invalid CSRF Token", "no CSRF Token"))
		}

		sid, ok := ctx.Get("sid")
		if !csrf.ValidateToken(token, sid.(string), m.logger) || !ok {
			m.logger.Errorf("CSRF Middleware get CSRF Header, Token: %s, Error: %s",
				token,
				"empty CSRF token",
			)
			utils.ErrorJSON(ctx, http.StatusForbidden, errors.NewRestError(http.StatusForbidden, "Invalid CSRF Token", "no CSRF Token"))
		}

		ctx.Next()
	}
}
