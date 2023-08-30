package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type RoleRoutes struct {
	logger  logger.Logger
	router  router.Router
	handler handlers.RoleHandler
	// authMiddleware       middlewares.AuthMiddleware
	PaginationMiddleware *middlewares.PaginationMiddleware
}

func NewRoleRoutes(
	logger logger.Logger,
	router router.Router,
	handler handlers.RoleHandler,
	// authMiddleware middlewares.AuthMiddleware,
	pagination *middlewares.PaginationMiddleware,
) *RoleRoutes {
	return &RoleRoutes{
		logger:  logger,
		router:  router,
		handler: handler,
		// authMiddleware:       authMiddleware,
		PaginationMiddleware: pagination,
	}
}

func (r *RoleRoutes) Setup() {
	r.logger.Info("Setting up roles routes")

	api := r.router.Group("/api/auth")

	api.GET("/roles", r.PaginationMiddleware.Handle(), r.handler.Index)
	api.GET("/roles/:id", r.handler.Show)
	api.POST("/roles", r.handler.Create)
	api.PUT("/roles/:id", r.handler.Update)
	api.DELETE("/roles/:id", r.handler.Delete)
}
