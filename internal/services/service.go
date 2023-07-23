package services

import (
	authServices "github.com/fazanurfaizi/go-rest-template/internal/services/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(authServices.NewUserService),
)
