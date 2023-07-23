package bootstrap

import (
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/internal/repositories"
	"github.com/fazanurfaizi/go-rest-template/internal/services"
	"github.com/fazanurfaizi/go-rest-template/pkg"
	"github.com/fazanurfaizi/go-rest-template/pkg/core"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	core.Module,
	pkg.Module,
	middlewares.Module,
	repositories.Module,
	services.Module,
)
