package repository

import (
	"github.com/fazanurfaizi/go-rest-template/internal/models"
	"github.com/gin-gonic/gin"
)

type AuthRepository interface {
	Create(c *gin.Context, user *models.User) (*models.User, error)
}
