package bootstrap

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth"
	"go.uber.org/fx"
)

var RepositoryModule = fx.Options(
	auth.AuthRepository,
)
