package router

import (
	"net/http"

	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(config *config.Config, logger logger.Logger) Router {
	if config.Server.Mode != constants.Dev && config.Sentry.Dsn != "" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:         config.Sentry.Dsn,
			Environment: `rest-template` + config.Server.Mode,
		}); err != nil {
			logger.Infof("Sentry initialization failed: %v \n", err)
		}
	}

	// gin.DefaultWriter = logger.GetGinLogger()
	if config.Server.Mode == constants.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	httpRouter := gin.Default()
	httpRouter.MaxMultipartMemory = constants.MaxMultipartMemory
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	// attach sentry middleware
	httpRouter.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	httpRouter.GET("/health-check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "go rest template API Up and Running"})
	})

	return Router{
		httpRouter,
	}
}
