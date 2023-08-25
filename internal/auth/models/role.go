package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"column:name;type:varchar(255);not null" filter:"param:name;searchable;filterable"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Permissions []Permission `gorm:"many2many:auth.role_permissions;"`
}

func (r *Role) TableName() string {
	return "auth.roles"
}
