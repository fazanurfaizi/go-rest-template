package dto

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
)

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name,omitempty"`
	Email       string `json:"email,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	Country     string `json:"country,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Postcode    string `json:"postcode,omitempty"`
	Birthday    string `json:"birthday,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
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
