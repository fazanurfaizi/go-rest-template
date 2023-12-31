package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// JSON json response function
func JSON(ctx *gin.Context, statusCode int, data any) {
	ctx.IndentedJSON(statusCode, gin.H{"data": data})
}

// ErrorJSON json error response function
func ErrorJSON(ctx *gin.Context, statusCode int, err any) {
	ctx.IndentedJSON(statusCode, gin.H{"error": err})
}

// SuccessJSON json success response function
func SuccessJSON(ctx *gin.Context, statusCode int, message any, data any) {
	ctx.IndentedJSON(statusCode, gin.H{"message": message, "data": data})
}

// JSONWithPagination json response with pagination meta function
func JSONWithPagination(ctx *gin.Context, statusCode int, response map[string]any) {
	limit, _ := ctx.MustGet(constants.Limit).(int64)
	page, _ := ctx.MustGet(constants.Page).(int64)
	total, _ := response["total"].(int64)

	pagination := PaginationMeta{
		Page:  page,
		Limit: limit,
		Total: total,
	}

	ctx.JSON(
		statusCode,
		gin.H{
			"data":       response["data"],
			"pagination": pagination,
		},
	)
}

// ValidationErrorJSON json error validation response function
func ValidationErrorJSON(ctx *gin.Context, err error) {
	validationErrors := make(map[string]interface{})
	for _, e := range err.(validator.ValidationErrors) {
		validationErrors[strings.ToLower(e.Field())] = e.Error()
	}

	jsonErrors, _ := json.Marshal(validationErrors)
	r := make(map[string]interface{})
	_ = json.Unmarshal([]byte(jsonErrors), &r)

	ErrorJSON(ctx, http.StatusBadRequest, r)
}

type PaginationMeta struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}
