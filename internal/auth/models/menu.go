package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID           uint  `gorm:"primarykey"`
	ParentID     *uint `gorm:"foreignkey:ParentID;nullable;default:null"`
	MasterMenuID uint  `gorm:"foreignkey:MasterMenuRefer"`
	MenuItemID   uint  `gorm:"foreignkey:MenuItemRefer"`
	Order        int   `gorm:"int"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Parent       *Menu      `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Children     []Menu     `gorm:"-"`
	MasterMenu   MasterMenu `gorm:"foreignKey:MasterMenuID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MenuItem     MenuItem   `gorm:"foreignKey:MenuItemID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (m *Menu) TableName() string {
	return "auth.menus"
}
