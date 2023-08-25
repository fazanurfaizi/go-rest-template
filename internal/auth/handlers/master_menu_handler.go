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

type MasterMenuHandler interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type masterMenuHandler struct {
	service *services.MasterMenuService
	logger  logger.Logger
}

func NewMasterMenuHandler(service *services.MasterMenuService, logger logger.Logger) MasterMenuHandler {
	return &masterMenuHandler{
		service: service,
		logger:  logger,
	}
}

// Index implements MasterMenuHandler.
func (h *masterMenuHandler) Index(ctx *gin.Context) {
	data, total := h.service.FindAll(ctx)

	response := map[string]any{"data": data, "total": total}

	utils.JSONWithPagination(ctx, http.StatusOK, response)
}

// Show implements MasterMenuHandler.
func (h *masterMenuHandler) Show(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	data, err := h.service.FindById(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
	}

	utils.JSON(ctx, http.StatusOK, data)
}

// Create implements MasterMenuHandler.
func (h *masterMenuHandler) Create(ctx *gin.Context) {
	var request dto.CreateMasterMenuRequest

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

// Update implements MasterMenuHandler.
func (h *masterMenuHandler) Update(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	var request dto.UpdateMasterMenuRequest

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

// Delete implements MasterMenuHandler.
func (h *masterMenuHandler) Delete(ctx *gin.Context) {
	param := ctx.Param("id")
	id, _ := strconv.Atoi(param)

	err := h.service.Delete(uint(id))
	if err != nil {
		utils.ErrorJSON(ctx, err.Status(), err.Error())
		return
	}

	utils.SuccessJSON(ctx, http.StatusOK, "Permission deleted successfully", nil)
}
