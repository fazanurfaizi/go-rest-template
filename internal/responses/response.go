package responses

import (
	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/gin-gonic/gin"
)

// JSON json response function
func JSON(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, gin.H{"data": data})
}

// ErrorJSON json error response function
func ErrorJSON(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, gin.H{"error": data})
}

// SuccessJSON json success response function
func SuccessJSON(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, gin.H{"message": data})
}

// JSONWithPagination json response with pagination meta function
func JSONWithPagination(ctx *gin.Context, statusCode int, response map[string]any) {
	limit, _ := ctx.MustGet(constants.Limit).(int64)
	size, _ := ctx.MustGet(constants.Page).(int64)
	page, _ := response["page"].(int64)
	total, _ := response["total"].(int64)

	pagination := PaginationMeta{
		Page:    page,
		PerPage: size,
		Limit:   limit,
		Total:   total,
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
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
	Limit   int64 `json:"limit"`
	Total   int64 `json:"total"`
}
