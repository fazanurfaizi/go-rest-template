package authHandler

import (
	"net/http"

	"github.com/fazanurfaizi/go-rest-template/internal/responses"
	authServices "github.com/fazanurfaizi/go-rest-template/internal/services/auth"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Index(ctx *gin.Context)
}

type userHandler struct {
	service *authServices.UserService
	logger  logger.Logger
}

func NewUserHandler(service *authServices.UserService, logger logger.Logger) UserHandler {
	return &userHandler{
		service: service,
		logger:  logger,
	}
}

func (u *userHandler) Index(ctx *gin.Context) {
	users, total, err := u.service.FindAll(ctx)
	if err != nil {
		u.logger.Error(err)
	}

	response := map[string]any{"data": users, "total": total}

	responses.JSONWithPagination(ctx, http.StatusOK, response)
}
