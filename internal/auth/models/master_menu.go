package models

import (
	"time"

	"gorm.io/gorm"
)

type MasterMenu struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"column:name;type:varchar(255);not null" filter:"param:name;searchable;filterable"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Menus     []Menu `gorm:"many2many:auth.menus;"`
}

func (m *MasterMenu) TableName() string {
	return "auth.master_menus"
}
