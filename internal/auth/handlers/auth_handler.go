package handlers

import (
	"net/http"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/services"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
}

type authHandler struct {
	service *services.AuthService
	logger  logger.Logger
}

func NewAuthHandler(service *services.AuthService, logger logger.Logger) AuthHandler {
	return &authHandler{
		service: service,
		logger:  logger,
	}
}

// Login implements AuthHandler.
func (h *authHandler) Login(ctx *gin.Context) {
	var request dto.LoginRequest

	if err := ctx.BindJSON(&request); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	data, err := h.service.Login(request)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusOK, "Login successfully", data)
}
