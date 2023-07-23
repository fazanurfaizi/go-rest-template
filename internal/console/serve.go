package console

import (
	"time"

	"github.com/fazanurfaizi/go-rest-template/pkg/command"
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
)

type ServeCommand struct{}

func (s *ServeCommand) Short() string {
	return "serve application"
}

func (s *ServeCommand) Setup(cmd *cobra.Command) {}

func (s *ServeCommand) Run() command.CommandRunner {
	return func(
		config *config.Config,
		router router.Router,
		logger logger.Logger,
		database postgres.Database,
	) {
		logger.Info(`+-----------------------+`)
		logger.Info(`| GO REST TEMPLATE      |`)
		logger.Info(`+-----------------------+`)

		// Using time zone as specified in config file
		loc, _ := time.LoadLocation("Asia/Jakarta")
		time.Local = loc

		// middleware.Setup()
		// routes.Setup()
		// seeds.Setup()

		if config.Server.Mode != constants.Dev && config.Sentry.Dsn != "" {
			err := sentry.Init(sentry.ClientOptions{
				Dsn:              config.Sentry.Dsn,
				AttachStacktrace: true,
			})
			if err != nil {
				logger.Error("Sentry initialization failed")
				logger.Error(err.Error())
			}
		}
		logger.Info("Running server")
		if config.Server.Port == "" {
			if err := router.Run(); err != nil {
				logger.Fatal(err)
				return
			}
		} else {
			if err := router.Run(config.Server.Port); err != nil {
				logger.Fatal(err)
				return
			}
		}
	}
}

func NewServeCommand() *ServeCommand {
	return &ServeCommand{}
}
