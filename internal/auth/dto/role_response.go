package dto

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
)

type RoleResponse struct {
	ID          uint                 `json:"id"`
	Name        string               `json:"name"`
	CreatedAt   string               `json:"created_at"`
	UpdatedAt   string               `json:"updated_at"`
	Permissions []PermissionResponse `json:"permissions,omitempty"`
}

func MappingRoleResponse(m models.Role) RoleResponse {
	var permissions []PermissionResponse
	if len(m.Permissions) > 0 {
		for _, v := range m.Permissions {
			permissions = append(permissions, MappingPermissionResponse(v))
		}
	}

	return RoleResponse{
		ID:          m.ID,
		Name:        m.Name,
		CreatedAt:   formatter.FormatTime(m.CreatedAt, formatter.YYYYMMDDhhmmss),
		UpdatedAt:   formatter.FormatTime(m.UpdatedAt, formatter.YYYYMMDDhhmmss),
		Permissions: permissions,
	}
}
