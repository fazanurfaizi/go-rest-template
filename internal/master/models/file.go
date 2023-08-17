package models

import (
	"database/sql"
	"time"
)

type File struct {
	ID           uint `gorm:"primarykey"`
	Filename     string
	FileableType string
	FileableId   uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    sql.NullTime `gorm:"index"`
}
