package bootstrap

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth"
	"go.uber.org/fx"
)

var HandlerModule = fx.Options(
	auth.AuthHandler,
)
