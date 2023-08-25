package dto

type CreatePermissionRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type UpdatePermissionRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
}
