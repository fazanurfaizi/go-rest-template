package auth

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/routes"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/services"
	"go.uber.org/fx"
)

// Auth Repositories Dependency Injection
var AuthRepository = fx.Options(
	fx.Provide(repositories.NewUserRepository),
	fx.Provide(repositories.NewRoleRepository),
	fx.Provide(repositories.NewPermissionRepository),
	fx.Provide(repositories.NewMenuItemRepository),
	fx.Provide(repositories.NewMasterMenuRepository),
	fx.Provide(repositories.NewAuthRepository),
)

// Auth Services Dependency Injection
var AuthService = fx.Options(
	fx.Provide(services.NewUserService),
	fx.Provide(services.NewRoleService),
	fx.Provide(services.NewPermissionService),
	fx.Provide(services.NewMenuItemService),
	fx.Provide(services.NewMasterMenuService),
	fx.Provide(services.NewAuthService),
)

// Auth Handlers Dependency Injection
var AuthHandler = fx.Options(
	fx.Provide(handlers.NewUserHandler),
	fx.Provide(handlers.NewPermissionHandler),
	fx.Provide(handlers.NewRoleHandler),
	fx.Provide(handlers.NewMenuItemHandler),
	fx.Provide(handlers.NewMasterMenuHandler),
	fx.Provide(handlers.NewAuthHandler),
)

// Auth Routes Dependency Injection
var AuthRoute = fx.Options(
	fx.Provide(routes.NewUserRoutes),
	fx.Provide(routes.NewRoleRoutes),
	fx.Provide(routes.NewPermissionRoutes),
	fx.Provide(routes.NewMenuItemRoutes),
	fx.Provide(routes.NewMasterMenuRoutes),
	fx.Provide(routes.NewAuthRoutes),
)
