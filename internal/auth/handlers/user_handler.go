package handlers

import (
	"net/http"
	"strconv"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/services"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
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
	users, total := u.service.FindAll(ctx)

	response := map[string]any{"data": users, "total": total}

	utils.JSONWithPagination(ctx, http.StatusOK, response)
}

func (u *userHandler) Show(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	user, err := u.service.FindById(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusNotFound, err)
	}

	utils.JSON(ctx, http.StatusOK, user)
}

func (u *userHandler) Create(ctx *gin.Context) {
	var userRequest dto.CreateUserRequest

	if err := ctx.MustBindWith(&userRequest, binding.FormMultipart); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	file, _, _ := ctx.Request.FormFile("image")

	user, err := u.service.Create(userRequest, file)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusOK, "User created successfully", user)
}
