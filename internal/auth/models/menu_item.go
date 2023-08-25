package models

import (
	"time"

	"gorm.io/gorm"
)

type MenuItem struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"column:name;type:varchar(255);not null" filter:"param:name;searchable;filterable"`
	Slug      string `gorm:"column:slug;type:varchar(255);not null"`
	Icon      string `gorm:"column:icon;type:varchar(255);not null"`
	Path      string `gorm:"column:path;type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (m *MenuItem) TableName() string {
	return "auth.menu_items"
}
