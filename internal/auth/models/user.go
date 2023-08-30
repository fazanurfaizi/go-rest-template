package models

import (
	"strings"
	"time"

	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"gorm.io/gorm"
)

type User struct {
	ID          uint      `gorm:"primarykey"`
	Name        string    `gorm:"column:name;type:varchar(255);not null" filter:"param:name;searchable;filterable"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;no null"`
	Password    string    `gorm:"type:varchar(255)"`
	Avatar      string    `gorm:"type:varchar(255)"`
	PhoneNumber string    `gorm:"type:varchar(255)"`
	Address     string    `gorm:"type:text"`
	City        string    `gorm:"type:varchar(255)"`
	Country     string    `gorm:"type:varchar(255)"`
	Gender      string    `gorm:"type:varchar(255)"`
	Postcode    string    `gorm:"type:varchar(10)"`
	Birthday    time.Time `gorm:"type:time"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (u *User) TableName() string {
	return "auth.users"
}

// Hash user password with bcrypt
func (u *User) HashPassword() error {
	hashedPassword, err := utils.GenerateHash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

// Compare user password and payload
func (u *User) ComparePassword(password string) (bool, error) {
	valid, err := utils.ValidateHash(u.Password, password)
	if err != nil {
		return false, err
	}
	return valid, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	hashedPassword, err := utils.GenerateHash(strings.TrimSpace(u.Password))
	u.Password = hashedPassword
	return err
}
