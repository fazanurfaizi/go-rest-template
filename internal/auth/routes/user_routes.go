package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type UserRoutes struct {
	logger                logger.Logger
	router                router.Router
	handler               handlers.UserHandler
	authMiddleware        *middlewares.AuthMiddleware
	PaginationMiddleware  *middlewares.PaginationMiddleware
	transactionMiddleware *middlewares.DBTransactionMiddleware
}

func NewUserRoutes(
	logger logger.Logger,
	router router.Router,
	handler handlers.UserHandler,
	authMiddleware *middlewares.AuthMiddleware,
	pagination *middlewares.PaginationMiddleware,
	transactionMiddleware *middlewares.DBTransactionMiddleware,
) *UserRoutes {
	return &UserRoutes{
		logger:                logger,
		router:                router,
		handler:               handler,
		authMiddleware:        authMiddleware,
		PaginationMiddleware:  pagination,
		transactionMiddleware: transactionMiddleware,
	}
}

func (r *UserRoutes) Setup() {
	r.logger.Info("Setting up user routes")

	api := r.router.Group("/api/auth").Use(r.authMiddleware.Handle())

	r.router.MaxMultipartMemory = 8 << 20
	api.GET("/users", r.PaginationMiddleware.Handle(), r.handler.Index)
	api.GET("/users/:id", r.handler.Show)
	api.POST("/users", r.transactionMiddleware.Handle(), r.handler.Create)
	api.PUT("/users/:id", r.transactionMiddleware.Handle(), r.handler.Update)
	api.DELETE("/users/:id", r.handler.Delete)
}
