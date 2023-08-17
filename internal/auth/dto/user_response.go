package dto

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
)

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	City        string `json:"city"`
	Country     string `json:"country"`
	Gender      string `json:"gender"`
	Postcode    string `json:"postcode"`
	Birthday    string `json:"birthday"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func MappingUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Avatar:      user.Avatar,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		City:        user.City,
		Country:     user.Country,
		Gender:      user.Gender,
		Postcode:    user.Postcode,
		Birthday:    formatter.FormatTime(user.Birthday, formatter.YYYYMMDDhhmmss),
		CreatedAt:   formatter.FormatTime(user.CreatedAt, formatter.YYYYMMDDhhmmss),
		UpdatedAt:   formatter.FormatTime(user.UpdatedAt, formatter.YYYYMMDDhhmmss),
	}
}
