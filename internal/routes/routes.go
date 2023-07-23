package routes

import (
	authRoutes "github.com/fazanurfaizi/go-rest-template/internal/routes/auth"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(authRoutes.NewUserRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(userRoutes *authRoutes.UserRoutes) Routes {
	return Routes{
		userRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
