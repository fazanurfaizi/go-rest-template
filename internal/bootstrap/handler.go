package bootstrap

import (
	authHandlers "github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"go.uber.org/fx"
)

var HandlerModule = fx.Options(
	fx.Provide(authHandlers.NewUserHandler),
	fx.Provide(authHandlers.NewRoleHandler),
	fx.Provide(authHandlers.NewPermissionHandler),
)
