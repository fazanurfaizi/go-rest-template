package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type RoleRoutes struct {
	logger      logger.Logger
	handler     router.Router
	roleHandler handlers.RoleHandler
	// authMiddleware       middlewares.AuthMiddleware
	PaginationMiddleware middlewares.PaginationMiddleware
	rateLimitMiddleware  middlewares.RateLimitMiddleware
}

func NewRoleRoutes(
	logger logger.Logger,
	handler router.Router,
	roleHandler handlers.RoleHandler,
	// authMiddleware middlewares.AuthMiddleware,
	pagination middlewares.PaginationMiddleware,
	rateLimitMiddleware middlewares.RateLimitMiddleware,
) *RoleRoutes {
	return &RoleRoutes{
		logger:      logger,
		handler:     handler,
		roleHandler: roleHandler,
		// authMiddleware:       authMiddleware,
		PaginationMiddleware: pagination,
		rateLimitMiddleware:  rateLimitMiddleware,
	}
}

func (r *RoleRoutes) Setup() {
	r.logger.Info("Setting up roles routes")

	api := r.handler.Group("/api").Use(r.rateLimitMiddleware.Handle())
	r.handler.MaxMultipartMemory = 8 << 20
	api.GET("/roles", r.PaginationMiddleware.Handle(), r.roleHandler.Index)
	// api.GET("/users/:id", r.userHandler.Show)
	// api.POST("/users", r.userHandler.Create)
	// api.PUT("/users/:id", r.userHandler.Update)
	// api.DELETE("/users/:id", r.userHandler.Delete)
}
