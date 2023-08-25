package bootstrap

import (
	authRepositories "github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"go.uber.org/fx"
)

var RepositoryModule = fx.Options(
	fx.Provide(authRepositories.NewUserRepository),
	fx.Provide(authRepositories.NewRoleRepository),
	fx.Provide(authRepositories.NewPermissionRepository),
)
