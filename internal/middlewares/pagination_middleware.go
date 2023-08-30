package middlewares

import (
	"strconv"

	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type PaginationMiddleware struct {
	logger logger.Logger
}

func NewPaginationMiddleware(logger logger.Logger) *PaginationMiddleware {
	return &PaginationMiddleware{logger: logger}
}

func (p *PaginationMiddleware) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p.logger.Info("Setting up pagination middleware")

		perPage, err := strconv.ParseInt(ctx.Query("page_size"), 10, 0)
		if err != nil {
			perPage = 10
		}

		page, err := strconv.ParseInt(ctx.Query("page"), 10, 0)
		if err != nil {
			page = 1
		}

		ctx.Set(constants.Limit, perPage)
		ctx.Set(constants.Page, page)

		ctx.Next()
	}
}
