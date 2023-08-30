package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type MenuItemRoutes struct {
	logger  logger.Logger
	router  router.Router
	handler handlers.MenuItemHandler
	// authMiddleware       middlewares.AuthMiddleware
	PaginationMiddleware *middlewares.PaginationMiddleware
}

func NewMenuItemRoutes(
	logger logger.Logger,
	router router.Router,
	handler handlers.MenuItemHandler,
	// authMiddleware middlewares.AuthMiddleware,
	pagination *middlewares.PaginationMiddleware,
) *MenuItemRoutes {
	return &MenuItemRoutes{
		logger:  logger,
		router:  router,
		handler: handler,
		// authMiddleware:       authMiddleware,
		PaginationMiddleware: pagination,
	}
}

func (r *MenuItemRoutes) Setup() {
	r.logger.Info("Setting up menu item routes")

	api := r.router.Group("/api/auth")

	api.GET("/menu-items", r.PaginationMiddleware.Handle(), r.handler.Index)
	api.GET("/menu-items/:id", r.handler.Show)
	api.POST("/menu-items", r.handler.Create)
	api.PUT("/menu-items/:id", r.handler.Update)
	api.DELETE("/menu-items/:id", r.handler.Delete)
}
