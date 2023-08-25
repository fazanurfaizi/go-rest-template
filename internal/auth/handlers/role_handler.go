package handlers

import (
	"net/http"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/services"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

type RoleHandler interface {
	Index(ctx *gin.Context)
	// show(ctx *gin.Context)
	// create(ctx *gin.Context)
	// update(ctx *gin.Context)
	// delete(ctx *gin.Context)
}

type roleHandler struct {
	service *services.RoleService
	logger  logger.Logger
}

func NewRoleHandler(service *services.RoleService, logger logger.Logger) RoleHandler {
	return &roleHandler{
		service: service,
		logger:  logger,
	}
}

func (u *roleHandler) Index(ctx *gin.Context) {
	users, total := u.service.FindAll(ctx)

	response := map[string]any{"data": users, "total": total}

	utils.JSONWithPagination(ctx, http.StatusOK, response)
}

// func (u *roleHandler) show(ctx *gin.Context) {
// 	param := ctx.Param("id")
// 	id, _ := strconv.Atoi(param)

// 	user, err := u.service.FindById(uint(id))
// 	if err != nil {
// 		utils.ErrorJSON(ctx, http.StatusNotFound, err)
// 	}

// 	utils.JSON(ctx, http.StatusOK, user)
// }

// func (u *roleHandler) create(ctx *gin.Context) {
// 	var request dto.CreateUserRequest

// 	if err := ctx.MustBindWith(&request, binding.FormMultipart); err != nil {
// 		utils.ValidationErrorJSON(ctx, err)
// 		return
// 	}

// 	file, _, _ := ctx.Request.FormFile("image")

// 	user, err := u.service.Create(request, file)
// 	if err != nil {
// 		utils.ErrorJSON(ctx, err.Status(), err.Error())
// 		return
// 	}

// 	utils.SuccessJSON(ctx, http.StatusCreated, "User created successfully", user)
// }

// func (u *roleHandler) update(ctx *gin.Context) {
// 	param := ctx.Param("id")
// 	id, _ := strconv.Atoi(param)

// 	var request dto.UpdateUserRequest

// 	if err := ctx.MustBindWith(&request, binding.FormMultipart); err != nil {
// 		utils.ValidationErrorJSON(ctx, err)
// 		return
// 	}

// 	file, _, _ := ctx.Request.FormFile("image")

// 	user, err := u.service.Update(uint(id), request, file)
// 	if err != nil {
// 		utils.ErrorJSON(ctx, err.Status(), err.Error())
// 		return
// 	}

// 	utils.SuccessJSON(ctx, http.StatusCreated, "User updated successfully", user)
// }

// func (u *roleHandler) dlete(ctx *gin.Context) {
// 	param := ctx.Param("id")
// 	id, _ := strconv.Atoi(param)
// 	err := u.service.Delete(uint(id))
// 	if err != nil {
// 		utils.ErrorJSON(ctx, err.Status(), err.Error())
// 		return
// 	}
// 	utils.SuccessJSON(ctx, http.StatusOK, "User deleted successfully", nil)
// }
