package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth"
	authRoutes "github.com/fazanurfaizi/go-rest-template/internal/auth/routes"
	"go.uber.org/fx"
)

var RouteModule = fx.Options(
	auth.AuthRoute,
	fx.Provide(NewRoutes),
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	userRoutes *authRoutes.UserRoutes,
	roleRoutes *authRoutes.RoleRoutes,
	permissionRoutes *authRoutes.PermissionRoutes,
	menuItemRoutes *authRoutes.MenuItemRoutes,
) Routes {
	return Routes{
		userRoutes,
		roleRoutes,
		permissionRoutes,
		menuItemRoutes,
	}
}

func (routes Routes) Setup() {
	for _, route := range routes {
		route.Setup()
	}
}
