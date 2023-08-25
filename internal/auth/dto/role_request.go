package dto

type CreateRoleRequest struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Permissions []uint `form:"permissions" json:"permissions" binding:"dive"`
}

type UpdateRoleRequest struct {
	Name        string `form:"name" json:"name" binding:"required"`
	Permissions []uint `form:"permissions" json:"permissions" binding:"dive"`
}
