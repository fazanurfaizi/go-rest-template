package dto

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
)

type MasterMenuResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Menus     []MenuResponse `json:"menus,omitempty"`
}

type MenuResponse struct {
	ID           uint             `json:"id"`
	ParentID     *uint            `json:"parent_id"`
	MasterMenuID uint             `json:"master_menu_id"`
	MenuItemID   uint             `json:"menu_item_id"`
	Order        int              `json:"order"`
	MenuItem     MenuItemResponse `json:"menu_item"`
	Children     []MenuResponse   `json:"children"`
	CreatedAt    string           `json:"created_at"`
	UpdatedAt    string           `json:"updated_at"`
}

func MappingMasterMenuResponse(m models.MasterMenu) MasterMenuResponse {
	return MasterMenuResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: formatter.FormatTime(m.CreatedAt, formatter.YYYYMMDDhhmmss),
		UpdatedAt: formatter.FormatTime(m.UpdatedAt, formatter.YYYYMMDDhhmmss),
	}
}
