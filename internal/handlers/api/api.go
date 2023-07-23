package api

import (
	authHandler "github.com/fazanurfaizi/go-rest-template/internal/handlers/api/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(authHandler.NewUserHandler),
)
