package dto

type CreateMasterMenuRequest struct {
	Name  string       `form:"name" json:"name" binding:"required"`
	Menus []CreateMenu `form:"menus" json:"menus" binding:"dive"`
}

type UpdateMasterMenuRequest struct {
	Name  string       `form:"name" json:"name" binding:"required"`
	Menus []CreateMenu `form:"menus" json:"menus" binding:"dive"`
}

type CreateMenu struct {
	ID         *uint        `form:"id" json:"id,omitempty" binding:"omitempty"`
	MenuItemID uint         `form:"menu_item_id" json:"menu_item_id" binding:"required"`
	Order      int          `form:"order" json:"order"`
	Children   []CreateMenu `form:"children" json:"children" binding:"dive"`
}
