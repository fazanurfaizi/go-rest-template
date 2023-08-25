package services

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type PermissionService struct {
	logger     logger.Logger
	repository repositories.PermissionRepository
}

func NewPermissionService(logger logger.Logger, repository repositories.PermissionRepository) *PermissionService {
	return &PermissionService{
		logger:     logger,
		repository: repository,
	}
}

func (s PermissionService) FindAll(ctx *gin.Context) ([]dto.PermissionResponse, int64) {
	var result []dto.PermissionResponse
	permissions, total := s.repository.FindAll(ctx)
	for _, permission := range permissions {
		result = append(result, dto.MappingPermissionResponse(permission))
	}

	return result, total
}

func (s PermissionService) FindById(id uint) (dto.PermissionResponse, errors.RestErr) {
	permission, err := s.repository.FindById(id)
	if err != nil {
		return dto.PermissionResponse{}, errors.NewNotFoundError(err.Error())
	}

	return dto.MappingPermissionResponse(permission), nil
}

func (s PermissionService) Create(request dto.CreatePermissionRequest) (dto.PermissionResponse, errors.RestErr) {
	permission := models.Permission{
		Name: request.Name,
	}

	createdPermission, err := s.repository.Create(&permission)
	if err != nil {
		return dto.PermissionResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingPermissionResponse(createdPermission), nil
}

func (s PermissionService) Update(id uint, request dto.UpdatePermissionRequest) (dto.PermissionResponse, errors.RestErr) {
	permission := models.Permission{
		Name: request.Name,
	}

	updatedPermission, err := s.repository.Update(id, &permission)
	if err != nil {
		return dto.PermissionResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingPermissionResponse(updatedPermission), nil
}

func (s PermissionService) Delete(id uint) errors.RestErr {
	err := s.repository.Delete(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
