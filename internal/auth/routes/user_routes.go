package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type UserRoutes struct {
	logger      logger.Logger
	handler     router.Router
	userHandler handlers.UserHandler
	// authMiddleware       middlewares.AuthMiddleware
	PaginationMiddleware middlewares.PaginationMiddleware
	rateLimitMiddleware  middlewares.RateLimitMiddleware
}

func NewUserRoutes(
	logger logger.Logger,
	handler router.Router,
	userHandler handlers.UserHandler,
	// authMiddleware middlewares.AuthMiddleware,
	pagination middlewares.PaginationMiddleware,
	rateLimitMiddleware middlewares.RateLimitMiddleware,
) *UserRoutes {
	return &UserRoutes{
		logger:      logger,
		handler:     handler,
		userHandler: userHandler,
		// authMiddleware:       authMiddleware,
		PaginationMiddleware: pagination,
		rateLimitMiddleware:  rateLimitMiddleware,
	}
}

func (r *UserRoutes) Setup() {
	r.logger.Info("Setting up user routes")

	api := r.handler.Group("/api").Use(r.rateLimitMiddleware.Handle())
	api.GET("/users", r.PaginationMiddleware.Handle(), r.userHandler.Index)
}
