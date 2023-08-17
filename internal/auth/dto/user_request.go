package dto

import "mime/multipart"

type CreateUserRequest struct {
	Name        string                `form:"name" json:"name" binding:"required"`
	Email       string                `form:"email" json:"email" binding:"required,email"`
	Password    string                `form:"password" json:"password" binding:"required"`
	Image       *multipart.FileHeader `form:"image" json:"image" binding:"omitempty"`
	PhoneNumber string                `form:"phone_number" json:"phone_number" binding:"required"`
	Address     string                `form:"address" json:"address" binding:"required"`
	City        string                `form:"city" json:"city" binding:"required"`
	Country     string                `form:"country" json:"country" binding:"required"`
	Gender      string                `form:"gender" json:"gender" binding:"required"`
	Postcode    string                `form:"post_code" json:"postcode" binding:"required"`
	Birthday    string                `form:"birthday" json:"birthday" binding:"required"`
	Avatar      string
}
