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

type MenuItemHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type menuItemHandler struct {
	service *services.MenuItemService
	logger  logger.Logger
}

func NewMenuItemHandler(service *services.MenuItemService, logger logger.Logger) MenuItemHandler {
	return &menuItemHandler{
		service: service,
		logger:  logger,
	}
}

// Index implements MenuItemHandler.
func (h *menuItemHandler) Index(ctx *gin.Context) {
	data, total := h.service.FindAll(ctx)

	response := map[string]any{"data": data, "total": total}

	utils.JSONWithPagination(ctx, http.StatusOK, response)
}

// Show implements MenuItemHandler.
func (h *menuItemHandler) Show(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	data, err := h.service.FindById(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
	}

	utils.JSON(ctx, http.StatusOK, data)
}

// Create implements MenuItemHandler.
func (h *menuItemHandler) Create(ctx *gin.Context) {
	var request dto.CreateMenuItemRequest

	if err := ctx.MustBindWith(&request, binding.JSON); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	data, err := h.service.Create(request)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Menu Item created successfully", data)
}

// Update implements MenuItemHandler.
func (h *menuItemHandler) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	var request dto.UpdateMenuItemRequest

	if err := ctx.MustBindWith(&request, binding.JSON); err != nil {
		utils.ValidationErrorJSON(ctx, err)
		return
	}

	data, err := h.service.Update(uint(id), request)
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusCreated, "Menu Item updated successfully", data)
}

// Delete implements MenuItemHandler.
func (h *menuItemHandler) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	err := h.service.Delete(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusOK, "Permission deleted successfully", nil)
}
