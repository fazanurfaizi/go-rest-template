package routes

import (
	authRoutes "github.com/fazanurfaizi/go-rest-template/internal/auth/routes"
	"go.uber.org/fx"
)

var RouteModule = fx.Options(
	fx.Provide(authRoutes.NewUserRoutes),
	fx.Provide(NewRoutes),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	userRoutes *authRoutes.UserRoutes,
) Routes {
	return Routes{
		userRoutes,
	}
}

func (routes Routes) Setup() {
	for _, route := range routes {
		route.Setup()
	}
}
