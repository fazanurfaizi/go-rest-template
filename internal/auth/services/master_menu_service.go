package services

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
)

type MasterMenuService struct {
	logger     logger.Logger
	repository repositories.MasterMenuRepository
}

func NewMasterMenuService(logger logger.Logger, repository repositories.MasterMenuRepository) *MasterMenuService {
	return &MasterMenuService{
		logger:     logger,
		repository: repository,
	}
}

func (s MasterMenuService) FindAll(ctx *gin.Context) ([]dto.MasterMenuResponse, int64) {
	var result = make([]dto.MasterMenuResponse, 0)
	masterMenus, total := s.repository.FindAll(ctx)
	for _, masterMenu := range masterMenus {
		result = append(result, dto.MappingMasterMenuResponse(masterMenu))
	}

	return result, total
}

func (s MasterMenuService) FindById(id uint) (dto.MasterMenuResponse, errors.RestErr) {
	masterMenu, err := s.repository.FindById(id)
	if err != nil {
		return dto.MasterMenuResponse{}, errors.NewNotFoundError(err.Error())
	}

	return dto.MappingMasterMenuResponse(masterMenu), nil
}

// Create implements MasterMenuService.
func (s *MasterMenuService) Create(request dto.CreateMasterMenuRequest) (dto.MasterMenuResponse, errors.RestErr) {
	var menus []models.Menu
	if len(request.Menus) > 0 {
		for i, v := range request.Menus {
			params := paramNormalizeMenu{
				menuItemID: v.MenuItemID,
				Menu:       v,
				parentID:   0,
				order:      i,
			}
			menu := normalizeMenu(params)

			if len(params.Menu.Children) > 0 {
				children := make([]models.Menu, 0)
				for k, child := range params.Menu.Children {
					childMenu := models.Menu{
						MenuItemID: child.MenuItemID,
						Order:      k,
					}

					if len(child.Children) > 0 {
						for _, c := range child.Children {
							childMenus := normalizeChildrenMenu(paramNormalizeChildrenMenu{
								menu:     c,
								children: c.Children,
								order:    k,
							})

							childMenu.Children = append(childMenu.Children, childMenus...)
						}
					}

					children = append(children, childMenu)
				}

				menu.Children = children
			}

			menus = append(menus, menu)
		}
	}

	masterMenu := models.MasterMenu{
		Name:  request.Name,
		Menus: menus,
	}

	createdMasterMenu, err := s.repository.Create(&masterMenu)
	if err != nil {
		return dto.MasterMenuResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingMasterMenuResponse(createdMasterMenu), nil
}

// Update implements MasterMenuService.
func (s *MasterMenuService) Update(id uint, request dto.UpdateMasterMenuRequest) (dto.MasterMenuResponse, errors.RestErr) {
	var menus []models.Menu
	if len(request.Menus) > 0 {
		for i, v := range request.Menus {
			params := paramNormalizeMenu{
				ID:         v.ID,
				menuItemID: v.MenuItemID,
				Menu:       v,
				parentID:   0,
				order:      i,
			}
			menu := normalizeMenu(params)

			if len(params.Menu.Children) > 0 {
				children := make([]models.Menu, 0)
				for k, child := range params.Menu.Children {
					childMenu := models.Menu{
						ID:         *child.ID,
						MenuItemID: child.MenuItemID,
						Order:      k,
					}

					if len(child.Children) > 0 {
						for _, c := range child.Children {
							childMenus := normalizeChildrenMenu(paramNormalizeChildrenMenu{
								menu:     c,
								children: c.Children,
								order:    k,
							})

							childMenu.Children = append(childMenu.Children, childMenus...)
						}
					}

					children = append(children, childMenu)
				}

				menu.Children = children
			}

			menus = append(menus, menu)
		}
	}

	masterMenu := models.MasterMenu{
		ID:    id,
		Name:  request.Name,
		Menus: menus,
	}

	updatedMasterMenu, err := s.repository.Update(id, &masterMenu)
	if err != nil {
		return dto.MasterMenuResponse{}, errors.NewInternalServerError(err.Error())
	}

	return dto.MappingMasterMenuResponse(updatedMasterMenu), nil
}

// Delete implements MasterMenuService.
func (s *MasterMenuService) Delete(id uint) errors.RestErr {
	err := s.repository.Delete(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}

type paramNormalizeMenu struct {
	ID         *uint
	menuItemID uint
	Menu       dto.CreateMenu
	parentID   uint
	order      int
}

func normalizeMenu(params paramNormalizeMenu) models.Menu {
	menu := models.Menu{
		MenuItemID: params.menuItemID,
		Order:      params.order,
		ParentID:   &params.parentID,
	}

	if params.ID != nil {
		menu.ID = *params.ID
	}

	if len(params.Menu.Children) > 0 {
		for k, child := range params.Menu.Children {
			childParams := paramNormalizeChildrenMenu{
				menu:  child,
				order: k,
			}
			menu.Children = normalizeChildrenMenu(childParams)
		}
	}

	return menu
}

type paramNormalizeChildrenMenu struct {
	menu     dto.CreateMenu
	children []dto.CreateMenu
	order    int
}

func normalizeChildrenMenu(params paramNormalizeChildrenMenu) []models.Menu {
	var results = make([]models.Menu, 0)
	menu := models.Menu{
		MenuItemID: params.menu.MenuItemID,
		Order:      params.order,
	}

	if params.menu.ID != nil {
		menu.ID = *params.menu.ID
	}

	for i, v := range params.children {
		childParams := paramNormalizeChildrenMenu{
			menu:     v,
			children: v.Children,
			order:    i,
		}
		menu.Children = normalizeChildrenMenu(childParams)
	}

	results = append(results, menu)

	return results
}
