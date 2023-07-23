package repositories

import (
	authRepositories "github.com/fazanurfaizi/go-rest-template/internal/repositories/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(authRepositories.NewUserRepository),
)
