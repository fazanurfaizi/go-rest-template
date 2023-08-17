package bootstrap

import (
	"github.com/fazanurfaizi/go-rest-template/database/faker"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/internal/routes"
	"github.com/fazanurfaizi/go-rest-template/pkg"
	"github.com/fazanurfaizi/go-rest-template/pkg/core"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	core.Module,
	pkg.Module,
	middlewares.Module,
	RepositoryModule,
	ServiceModule,
	HandlerModule,
	routes.RouteModule,
	faker.Module,
)
