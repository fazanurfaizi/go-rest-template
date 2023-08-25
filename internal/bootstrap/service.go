package bootstrap

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	auth.AuthService,
)
