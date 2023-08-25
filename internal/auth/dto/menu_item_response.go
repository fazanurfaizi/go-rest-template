package dto

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
)

type MenuItemResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func MappingMenuItemResponse(m models.MenuItem) MenuItemResponse {
	return MenuItemResponse{
		ID:        m.ID,
		Name:      m.Name,
		Slug:      m.Slug,
		Icon:      m.Icon,
		Path:      m.Path,
		CreatedAt: formatter.FormatTime(m.CreatedAt, formatter.YYYYMMDDhhmmss),
		UpdatedAt: formatter.FormatTime(m.UpdatedAt, formatter.YYYYMMDDhhmmss),
	}
}
