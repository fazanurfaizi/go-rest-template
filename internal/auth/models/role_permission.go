package models

import (
	"time"

	"gorm.io/gorm"
)

type RolePermission struct {
	RoleID       uint `gorm:"foreignKey:role_id"`
	PermissionID uint `gorm:"foreignKey:permission_id"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Role         Role       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Permission   Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (rp *RolePermission) TableName() string {
	return "auth.role_permissions"
}
