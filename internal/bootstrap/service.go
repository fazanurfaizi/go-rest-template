package bootstrap

import (
	authServices "github.com/fazanurfaizi/go-rest-template/internal/auth/services"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	fx.Provide(authServices.NewUserService),
	fx.Provide(authServices.NewAuthService),
	fx.Provide(authServices.NewRoleService),
	fx.Provide(authServices.NewPermissionService),
)
