package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"column:name;type:varchar(255);not null" filter:"param:name,searchable;filterable"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Deleted   gorm.DeletedAt
	Roles     []Role `gorm:"many2many:auth.role_permissions;"`
}

func (p *Permission) TableName() string {
	return "auth.permissions"
}
