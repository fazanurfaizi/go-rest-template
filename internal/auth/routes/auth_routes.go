package routes

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/handlers"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/router"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type AuthRoutes struct {
	logger  logger.Logger
	router  router.Router
	handler handlers.AuthHandler
	// authMiddleware       middlewares.AuthMiddleware
}

func NewAuthRoutes(
	logger logger.Logger,
	router router.Router,
	handler handlers.AuthHandler,
	// authMiddleware middlewares.AuthMiddleware,
) *AuthRoutes {
	return &AuthRoutes{
		logger:  logger,
		router:  router,
		handler: handler,
		// authMiddleware:       authMiddleware,
	}
}

func (r *AuthRoutes) Setup() {
	r.logger.Info("Setting up auth routes")

	api := r.router.Group("/api/auth")
	api.POST("/login", r.handler.Login)

}
