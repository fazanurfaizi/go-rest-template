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
	PaginationMiddleware  middlewares.PaginationMiddleware
	rateLimitMiddleware   middlewares.RateLimitMiddleware
	transactionMiddleware middlewares.DBTransactionMiddleware
}

func NewUserRoutes(
	logger logger.Logger,
	handler router.Router,
	userHandler handlers.UserHandler,
	// authMiddleware middlewares.AuthMiddleware,
	pagination middlewares.PaginationMiddleware,
	rateLimitMiddleware middlewares.RateLimitMiddleware,
	transactionMiddleware middlewares.DBTransactionMiddleware,
) *UserRoutes {
	return &UserRoutes{
		logger:      logger,
		handler:     handler,
		userHandler: userHandler,
		// authMiddleware:       authMiddleware,
		PaginationMiddleware:  pagination,
		rateLimitMiddleware:   rateLimitMiddleware,
		transactionMiddleware: transactionMiddleware,
	}
}

func (r *UserRoutes) Setup() {
	r.logger.Info("Setting up user routes")

	api := r.handler.Group("/api").
		Use(r.rateLimitMiddleware.Handle()).
		Use(r.transactionMiddleware.Handle())

	r.handler.MaxMultipartMemory = 8 << 20
	api.GET("/users", r.PaginationMiddleware.Handle(), r.userHandler.Index)
	api.GET("/users/:id", r.userHandler.Show)
	api.POST("/users", r.userHandler.Create)
	api.PUT("/users/:id", r.userHandler.Update)
	api.DELETE("/users/:id", r.userHandler.Delete)
}
