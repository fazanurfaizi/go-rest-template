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

type RoleHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type roleHandler struct {
	service services.RoleService
	logger  logger.Logger
}

func NewRoleHandler(service services.RoleService, logger logger.Logger) RoleHandler {
	return &roleHandler{
		service: service,
		logger:  logger,
	}
}

func (h *roleHandler) Index(ctx *gin.Context) {
	users, total := h.service.FindAll(ctx)

	response := map[string]any{"data": users, "total": total}

	utils.JSONWithPagination(ctx, http.StatusOK, response)
}

func (h *roleHandler) Show(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	user, err := h.service.FindById(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, http.StatusNotFound, err)
	}

	utils.JSON(ctx, http.StatusOK, user)
}

func (h *roleHandler) Create(ctx *gin.Context) {
	var request dto.CreateRoleRequest

	if err := ctx.MustBindWith(&request, binding.JSON); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	role, err := h.service.Create(request)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Role created successfully", role)
}

func (h *roleHandler) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	var request dto.UpdateRoleRequest

	if err := ctx.MustBindWith(&request, binding.JSON); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	role, err := h.service.Update(uint(id), request)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Role updated successfully", role)
}

func (h *roleHandler) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	err := h.service.Delete(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusOK, "Role deleted successfully", nil)
}
