package services

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type RoleService interface {
	FindAll(*gin.Context) ([]dto.RoleResponse, int64)
	FindById(uint) (dto.RoleResponse, errors.RestErr)
	Create(dto.CreateRoleRequest) (dto.RoleResponse, errors.RestErr)
	Update(uint, dto.UpdateRoleRequest) (dto.RoleResponse, errors.RestErr)
	Delete(uint) errors.RestErr
}

type roleService struct {
	logger     logger.Logger
	repository repositories.RoleRepository
}

func NewRoleService(logger logger.Logger, repository repositories.RoleRepository) RoleService {
	return &roleService{
		logger:     logger,
		repository: repository,
	}
}

func (s *roleService) FindAll(ctx *gin.Context) ([]dto.RoleResponse, int64) {
	var result = make([]dto.RoleResponse, 0)
	roles, total := s.repository.FindAll(ctx)
	for _, role := range roles {
		result = append(result, dto.MappingRoleResponse(role))
	}

	return result, total
}

func (s *roleService) FindById(id uint) (dto.RoleResponse, errors.RestErr) {
	role, err := s.repository.FindById(id)
	if err != nil {
		return dto.RoleResponse{}, errors.NewNotFoundError(err.Error())
	}

	return dto.MappingRoleResponse(role), nil
}

func (s *roleService) Create(request dto.CreateRoleRequest) (dto.RoleResponse, errors.RestErr) {
	var permissions []models.Permission
	if len(request.Permissions) > 0 {
		for _, v := range request.Permissions {
			permission := models.Permission{ID: v}
			permissions = append(permissions, permission)
		}
	}

	role := models.Role{
		Name:        request.Name,
		Permissions: permissions,
	}

	createdRole, err := s.repository.Create(&role)
	if err != nil {
		return dto.RoleResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingRoleResponse(createdRole), nil
}

func (s *roleService) Update(id uint, request dto.UpdateRoleRequest) (dto.RoleResponse, errors.RestErr) {
	var permissions []models.Permission
	if len(request.Permissions) > 0 {
		for _, v := range request.Permissions {
			permission := models.Permission{ID: v}
			permissions = append(permissions, permission)
		}
	}

	role := models.Role{
		Name:        request.Name,
		Permissions: permissions,
	}

	updatedRole, err := s.repository.Update(id, &role)
	if err != nil {
		return dto.RoleResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingRoleResponse(updatedRole), nil
}

// Delete implements RoleService.
func (s *roleService) Delete(id uint) errors.RestErr {
	err := s.repository.Delete(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
