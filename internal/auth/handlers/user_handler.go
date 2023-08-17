package handlers

import (
	"net/http"
	"strconv"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/services"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
}

type userHandler struct {
	service *services.UserService
	logger  logger.Logger
}

func NewUserHandler(service *services.UserService, logger logger.Logger) UserHandler {
	return &userHandler{
		service: service,
		logger:  logger,
	}
}

func (u *userHandler) Index(ctx *gin.Context) {
	var usersResponse []dto.UserResponse
	users, total := u.service.FindAll(ctx)
	for _, user := range users {
		usersResponse = append(usersResponse, dto.MappingUserResponse(user))
	}

	response := map[string]any{"data": usersResponse, "total": total}

	utils.JSONWithPagination(ctx, http.StatusOK, response)
}

func (u *userHandler) Show(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	user := u.service.FindById(uint(id))
	userResponse := dto.MappingUserResponse(user)

	utils.JSON(ctx, http.StatusOK, userResponse)
}
