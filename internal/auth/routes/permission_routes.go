package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type PermissionRoutes struct {
	logger  logger.Logger
	router  router.Router
	handler handlers.PermissionHandler
	// authMiddleware       middlewares.AuthMiddleware
	PaginationMiddleware middlewares.PaginationMiddleware
}

func NewPermissionRoutes(
	logger logger.Logger,
	router router.Router,
	handler handlers.PermissionHandler,
	// authMiddleware middlewares.AuthMiddleware,
	pagination middlewares.PaginationMiddleware,
) *PermissionRoutes {
	return &PermissionRoutes{
		logger:  logger,
		router:  router,
		handler: handler,
		// authMiddleware:       authMiddleware,
		PaginationMiddleware: pagination,
	}
}

func (r *PermissionRoutes) Setup() {
	r.logger.Info("Setting up permission routes")

	api := r.router.Group("/api")

	api.GET("/permissions", r.PaginationMiddleware.Handle(), r.handler.Index)
	api.GET("/permissions/:id", r.handler.Show)
	api.POST("/permissions", r.handler.Create)
	api.PUT("/permissions/:id", r.handler.Update)
	api.DELETE("/permissions/:id", r.handler.Delete)
}
