package dto

import "github.com/fazanurfaizi/go-rest-template/internal/auth/models"

type LoginRequest struct {
	Email    string `json:"email" binding:"omitempty,lte=60,email"`
	Password string `json:"password,omitempty" binding:"required,gte=6"`
}

type LoginResponse struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

type RegisterRequesst struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"omitempty,lte=60,email"`
	Password string `json:"password,omitempty" binding:"required,gte=6"`
}

type RegisterResponse struct {
	VerificationToken string `json:"verification_token"`
}
