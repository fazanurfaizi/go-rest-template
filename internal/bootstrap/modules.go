package bootstrap

import (
	"github.com/fazanurfaizi/go-rest-template/pkg"
	"github.com/fazanurfaizi/go-rest-template/pkg/core"
	"github.com/fazanurfaizi/go-rest-template/pkg/middlewares"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	core.Module,
	pkg.Module,
	RepositoryModule,
	ServiceModule,
	middlewares.Module,
)
