package pkg

import (
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(config.NewConfig),
	fx.Provide(logger.GetLogger),
)
