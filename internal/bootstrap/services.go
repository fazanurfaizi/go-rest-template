package bootstrap

import (
	AuthService "github.com/fazanurfaizi/go-rest-template/internal/auth/services"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	AuthService.Module,
)
