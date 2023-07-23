package bootstrap

import (
	"github.com/fazanurfaizi/go-rest-template/pkg"
	"github.com/fazanurfaizi/go-rest-template/pkg/core"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	core.Module,
	pkg.Module,
)
