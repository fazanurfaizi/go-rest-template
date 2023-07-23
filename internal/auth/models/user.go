package models

import (
	"database/sql"
	"strings"
	"time"

	"github.com/fazanurfaizi/go-rest-template/internal/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	models.BaseModel
	Name        string    `gorm:"type:varchar(255);not null" json:"name" redis:"name" validate:"required,lte=30"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;no null" json:"email"`
	Password    string    `gorm:"type:varchar(255)" json:"password,omitempty" redis:"password" validate:"omitempty,required,gte=6"`
	Avatar      string    `gorm:"type:varchar(255)" json:"avatar,omitempty" redis:"avatar" validate:"omitempty,lte=512,url"`
	PhoneNumber string    `gorm:"type:varchar(255)" json:"phone_number,omitempty" redis:"phone_number" validate:"omitempty,lte=20"`
	Address     string    `gorm:"type:text" json:"address,omitempty" redis:"address" validate:"omitempty,lte=250"`
	City        string    `gorm:"type:varchar(255)" json:"city,omitempty" redis:"city" validate:"omitempty,lte=24"`
	Country     string    `gorm:"type:varchar(255)" json:"country,omitempty" redis:"country" validate:"omitempty,lte=24"`
	Gender      string    `gorm:"type:varchar(255)" json:"gender,omitempty" redis:"gender" validate:"omitempty,lte=10"`
	Postcode    string    `gorm:"type:varchar(10)" json:"postcode,omitempty" redis:"postcode" validate:"omitempty"`
	Birthday    time.Time `gorm:"type:time" json:"birthday,omitempty" redis:"birthday" validate:"omitempty,lte=10"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime `gorm:"index"`
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
	valid, err := utils.ValidateHash(password, u.Password)
	if err != nil {
		return false, err
	}
	return valid, nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewRandom()
	u.ID = models.BinaryUUID(id)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))
	u.Password = strings.TrimSpace(u.Password)

	if err := u.HashPassword(); err != nil {
		return err
	}

	return err
}
