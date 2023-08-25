package dto

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
)

type MasterMenuResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	// Permissions []PermissionResponse `json:"permissions,omitempty"`
}

func MappingMasterMenuResponse(m models.MasterMenu) MasterMenuResponse {
	// var permissions []PermissionResponse
	// if len(m.Permissions) > 0 {
	// 	for _, v := range m.Permissions {
	// 		permissions = append(permissions, MappingPermissionResponse(v))
	// 	}
	// }

	return MasterMenuResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: formatter.FormatTime(m.CreatedAt, formatter.YYYYMMDDhhmmss),
		UpdatedAt: formatter.FormatTime(m.UpdatedAt, formatter.YYYYMMDDhhmmss),
		// Permissions: permissions,
	}
}
