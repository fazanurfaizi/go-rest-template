package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/internal/middlewares"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type MasterMenuRoutes struct {
	logger  logger.Logger
	router  router.Router
	handler handlers.MasterMenuHandler
	// authMiddleware       middlewares.AuthMiddleware
	PaginationMiddleware  middlewares.PaginationMiddleware
	transactionMiddleware middlewares.DBTransactionMiddleware
}

func NewMasterMenuRoutes(
	logger logger.Logger,
	router router.Router,
	handler handlers.MasterMenuHandler,
	// authMiddleware middlewares.AuthMiddleware,
	pagination middlewares.PaginationMiddleware,
	transactionMiddleware middlewares.DBTransactionMiddleware,
) *MasterMenuRoutes {
	return &MasterMenuRoutes{
		logger:  logger,
		router:  router,
		handler: handler,
		// authMiddleware:       authMiddleware,
		PaginationMiddleware:  pagination,
		transactionMiddleware: transactionMiddleware,
	}
}

func (r *MasterMenuRoutes) Setup() {
	r.logger.Info("Setting up master menu routes")

	api := r.router.Group("/api/auth")

	api.GET("/master-menus", r.PaginationMiddleware.Handle(), r.handler.Index)
	api.GET("/master-menus/:id", r.handler.Show)
	api.POST("/master-menus", r.transactionMiddleware.Handle(), r.handler.Create)
	api.PUT("/master-menus/:id", r.transactionMiddleware.Handle(), r.handler.Update)
	api.DELETE("/master-menus/:id", r.handler.Delete)
}
