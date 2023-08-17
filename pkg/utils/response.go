package utils

import (
	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/gin-gonic/gin"
)

// JSON json response function
func JSON(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, gin.H{"data": data})
}

// ErrorJSON json error response function
func ErrorJSON(ctx *gin.Context, statusCode int, err any) {
	ctx.JSON(statusCode, gin.H{"error": err})
}

// SuccessJSON json success response function
func SuccessJSON(ctx *gin.Context, statusCode int, message any) {
	ctx.JSON(statusCode, gin.H{"message": message})
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

type PaginationMeta struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
	Total int64 `json:"total"`
}
