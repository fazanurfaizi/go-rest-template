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
)

// Auth Services Dependency Injection
var AuthService = fx.Options(
	fx.Provide(services.NewAuthService),
	fx.Provide(services.NewUserService),
	fx.Provide(services.NewRoleService),
	fx.Provide(services.NewPermissionService),
)

// Auth Handlers Dependency Injection
var AuthHandler = fx.Options(
	fx.Provide(handlers.NewUserHandler),
	fx.Provide(handlers.NewPermissionHandler),
	fx.Provide(handlers.NewRoleHandler),
)

// Auth Routes Dependency Injection
var AuthRoute = fx.Options(
	fx.Provide(routes.NewUserRoutes),
	fx.Provide(routes.NewRoleRoutes),
	fx.Provide(routes.NewPermissionRoutes),
)
