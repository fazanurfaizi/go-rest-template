package bootstrap

import (
	authRepository "github.com/fazanurfaizi/go-rest-template/internal/auth/repository"
	"go.uber.org/fx"
)

var RepositoryModule = fx.Options(
	authRepository.Module,
)
