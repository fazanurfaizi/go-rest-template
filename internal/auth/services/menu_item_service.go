package services

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type MenuItemService interface {
	FindAll(*gin.Context) ([]dto.MenuItemResponse, int64)
	FindById(uint) (dto.MenuItemResponse, errors.RestErr)
	Create(dto.CreateMenuItemRequest) (dto.MenuItemResponse, errors.RestErr)
	Update(uint, dto.UpdateMenuItemRequest) (dto.MenuItemResponse, errors.RestErr)
	Delete(uint) errors.RestErr
}

type menuItemService struct {
	logger     logger.Logger
	repository repositories.MenuItemRepository
}

func NewMenuItemService(logger logger.Logger, repository repositories.MenuItemRepository) MenuItemService {
	return &menuItemService{
		logger:     logger,
		repository: repository,
	}
}

func (s *menuItemService) FindAll(ctx *gin.Context) ([]dto.MenuItemResponse, int64) {
	var result = make([]dto.MenuItemResponse, 0)
	menuItems, total := s.repository.FindAll(ctx)
	for _, menuItem := range menuItems {
		result = append(result, dto.MappingMenuItemResponse(menuItem))
	}

	return result, total
}

func (s *menuItemService) FindById(id uint) (dto.MenuItemResponse, errors.RestErr) {
	menuItem, err := s.repository.FindById(id)
	if err != nil {
		return dto.MenuItemResponse{}, errors.NewNotFoundError(err.Error())
	}

	return dto.MappingMenuItemResponse(menuItem), nil
}

func (s *menuItemService) Create(request dto.CreateMenuItemRequest) (dto.MenuItemResponse, errors.RestErr) {
	menuItem := models.MenuItem{
		Name: request.Name,
		Slug: request.Slug,
		Icon: request.Icon,
		Path: request.Path,
	}

	createdMenuItem, err := s.repository.Create(&menuItem)
	if err != nil {
		return dto.MenuItemResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingMenuItemResponse(createdMenuItem), nil
}

func (s *menuItemService) Update(id uint, request dto.UpdateMenuItemRequest) (dto.MenuItemResponse, errors.RestErr) {
	menuItem := models.MenuItem{
		Name: request.Name,
		Slug: request.Slug,
		Icon: request.Icon,
		Path: request.Path,
	}

	updatedMenuItem, err := s.repository.Update(id, &menuItem)
	if err != nil {
		return dto.MenuItemResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingMenuItemResponse(updatedMenuItem), nil
}

func (s *menuItemService) Delete(id uint) errors.RestErr {
	err := s.repository.Delete(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
