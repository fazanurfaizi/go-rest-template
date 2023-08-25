package dto

type CreateMenuItemRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
	Slug string `form:"slug" json:"slug" binding:"required"`
	Icon string `form:"icon" json:"icon" binding:"required"`
	Path string `form:"path" json:"path" binding:"required"`
}

type UpdateMenuItemRequest struct {
	Name string `form:"name" json:"name" binding:"required"`
	Slug string `form:"slug" json:"slug" binding:"required"`
	Icon string `form:"icon" json:"icon" binding:"required"`
	Path string `form:"path" json:"path" binding:"required"`
}
