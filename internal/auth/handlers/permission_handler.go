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

type PermissionHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type permissionHandler struct {
	service *services.PermissionService
	logger  logger.Logger
}

func NewPermissionHandler(service *services.PermissionService, logger logger.Logger) PermissionHandler {
	return &permissionHandler{
		service: service,
		logger:  logger,
	}
}

// Index implements PermissionHandler.
func (h *permissionHandler) Index(ctx *gin.Context) {
	permissions, total := h.service.FindAll(ctx)

	response := map[string]any{"data": permissions, "total": total}

	utils.JSONWithPagination(ctx, http.StatusOK, response)
}

// Show implements PermissionHandler.
func (h *permissionHandler) Show(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	permission, err := h.service.FindById(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
	}

	utils.JSON(ctx, http.StatusOK, permission)
}

// Create implements PermissionHandler.
func (h *permissionHandler) Create(ctx *gin.Context) {
	var request dto.CreatePermissionRequest

	if err := ctx.MustBindWith(&request, binding.JSON); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	permission, err := h.service.Create(request)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Permission created successfully", permission)
}

// Update implements PermissionHandler.
func (h *permissionHandler) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	var request dto.UpdatePermissionRequest

	if err := ctx.MustBindWith(&request, binding.JSON); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	permission, err := h.service.Update(uint(id), request)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Permission updated successfully", permission)
}

// Delete implements PermissionHandler.
func (h *permissionHandler) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	err := h.service.Delete(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusOK, "Permission deleted successfully", nil)
}
