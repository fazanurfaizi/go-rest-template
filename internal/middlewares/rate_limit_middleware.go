package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// Global store
// using in-memory store with goroutine which clears expired keys.
var store = memory.NewStore()

type RateLimitOption struct {
	period time.Duration
	limit  int64
}

type Option func(*RateLimitOption)

type RateLimitMiddleware struct {
	logger logger.Logger
	option RateLimitOption
}

func NewRateLimitMiddleware(logger logger.Logger) RateLimitMiddleware {
	return RateLimitMiddleware{
		logger: logger,
		option: RateLimitOption{
			period: constants.RateLimitPeriod,
			limit:  constants.RateLimitRequests,
		},
	}
}

func (lm RateLimitMiddleware) Handle(options ...Option) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.ClientIP()

		lm.logger.Info("Setting up rate limit middleware")

		opt := RateLimitOption{
			period: lm.option.period,
			limit:  lm.option.limit,
		}

		for _, o := range options {
			o(&opt)
		}

		rate := limiter.Rate{
			Limit:  opt.limit,
			Period: opt.period,
		}

		instance := limiter.New(store, rate)

		context, err := instance.Get(ctx, ctx.FullPath()+"&&"+key)
		if err != nil {
			lm.logger.Panic(err.Error())
		}

		ctx.Set(constants.RateLimit, instance)

		// Setting header
		ctx.Header("X-RateLimit-Limit", strconv.FormatInt(context.Limit, 10))
		ctx.Header("X-RateLimit-Remaining", strconv.FormatInt(context.Remaining, 10))
		ctx.Header("X-RateLimit-Rest", strconv.FormatInt(context.Reset, 10))

		// Limit exceeded
		if context.Reached {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"message": "Rate limit exceed",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func WithOption(period time.Duration, limit int64) Option {
	return func(rlo *RateLimitOption) {
		rlo.period = period
		rlo.limit = limit
	}
}
