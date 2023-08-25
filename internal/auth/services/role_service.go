package services

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type RoleService struct {
	logger     logger.Logger
	repository repositories.RoleRepository
}

func NewRoleService(logger logger.Logger, repository repositories.RoleRepository) *RoleService {
	return &RoleService{
		logger:     logger,
		repository: repository,
	}
}

func (s RoleService) FindAll(ctx *gin.Context) ([]models.Role, int64) {
	roles, total := s.repository.FindAll(ctx)

	return roles, total
}
