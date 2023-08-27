package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDBTransactionMiddleware),
	fx.Provide(NewPaginationMiddleware),
	fx.Provide(NewRateLimitMiddleware),
	fx.Provide(NewJsonMiddleware),
	fx.Provide(NewCsrfMiddleware),
	fx.Provide(NewMiddlewares),
)

type IMiddleware interface {
	Handle() gin.HandlerFunc
}

type Middlewares []IMiddleware

func NewMiddlewares(
	jsonMiddleware *JsonMiddleware,
	rateLimitMiddleware *RateLimitMiddleware,
	csrfMiddleware *CsrfMiddleware,
) Middlewares {
	return Middlewares{
		jsonMiddleware,
		rateLimitMiddleware,
		csrfMiddleware,
	}
}

func (m Middlewares) Handle() {
	for _, middleware := range m {
		middleware.Handle()
	}
}
