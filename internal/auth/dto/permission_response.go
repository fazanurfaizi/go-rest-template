package dto

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
)

type PermissionResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func MappingPermissionResponse(m models.Permission) PermissionResponse {
	return PermissionResponse{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: formatter.FormatTime(m.CreatedAt, formatter.YYYYMMDDhhmmss),
		UpdatedAt: formatter.FormatTime(m.UpdatedAt, formatter.YYYYMMDDhhmmss),
	}
}
