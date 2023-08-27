package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres/filter"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type MasterMenuRepository struct {
	postgres.Database
	logger logger.Logger
}

func NewMasterMenuRepository(db postgres.Database, logger logger.Logger) MasterMenuRepository {
	return MasterMenuRepository{db, logger}
}

func (r MasterMenuRepository) FindAll(ctx *gin.Context) ([]models.MasterMenu, int64) {
	var masterMenus []models.MasterMenu
	var count int64

	r.Model(&models.MasterMenu{}).
		Scopes(filter.FilterByQuery(ctx, filter.ALL)).
		Find(&masterMenus).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return masterMenus, count
}

func (r MasterMenuRepository) FindById(id uint) (masterMenu models.MasterMenu, err error) {
	err = r.Model(&masterMenu).
		Where("id = ?", id).
		First(&masterMenu).
		Error

	if err != nil {
		return models.MasterMenu{}, err
	}

	var menuResult = make([]models.Menu, 0)
	err = r.DB.
		Model(&models.Menu{}).
		Preload("MenuItem").
		Where("master_menu_id = ?", id).
		Order("parent_id").
		Find(&menuResult).
		Error

	if err != nil {
		return models.MasterMenu{}, err
	}

	var menus = make([]models.Menu, 0)
	for _, menu := range menuResult {
		if menu.ParentID == nil {
			childMenus := normalizeMenus(&menu, menuResult)
			menu.Children = childMenus
			menus = append(menus, menu)
		}
	}

	masterMenu.Menus = menus

	return masterMenu, nil
}

func (r MasterMenuRepository) Create(masterMenu *models.MasterMenu) (models.MasterMenu, error) {
	createdMasterMenu := models.MasterMenu{Name: masterMenu.Name}

	err := r.Model(&createdMasterMenu).Clauses(clause.Returning{}).Select("id", "name").Create(&createdMasterMenu).Error

	if err != nil {
		return models.MasterMenu{}, err
	}

	if len(masterMenu.Menus) > 0 {
		for i, v := range masterMenu.Menus {
			menu, _ := saveMenu(&r, saveMenuParams{
				masterMenu: &createdMasterMenu,
				menu:       v,
				parentID:   nil,
				order:      i,
			})

			if len(v.Children) > 0 {
				for k, child := range v.Children {
					val := models.Menu{
						ParentID:     &menu.ID,
						MasterMenuID: createdMasterMenu.ID,
						MenuItemID:   child.MenuItemID,
						Order:        k,
					}

					childModel, _ := saveMenu(&r, saveMenuParams{
						masterMenu: &createdMasterMenu,
						menu:       val,
						parentID:   &menu.ID,
						order:      k,
					})

					if len(child.Children) > 0 {
						for _, c := range child.Children {
							saveChildrenMenu(&r, saveChildrenMenuParams{
								masterMenu: &createdMasterMenu,
								menu:       c,
								parentID:   &childModel.ID,
								menus:      c.Children,
							})
						}
					}
				}
			}
		}
	}

	return createdMasterMenu, nil
}

func (r MasterMenuRepository) Update(id uint, masterMenu *models.MasterMenu) (models.MasterMenu, error) {
	var updateMasterMenu = models.MasterMenu{}
	err := r.Model(&updateMasterMenu).
		Where("id = ?", id).
		Updates(&masterMenu).
		Error

	if err != nil {
		return models.MasterMenu{}, err
	}

	if err := r.DB.Model(&models.Menu{}).Delete(&models.Menu{}, "master_menu_id = ?", id).Error; err != nil {
		return models.MasterMenu{}, err
	}

	if len(masterMenu.Menus) > 0 {
		for i, v := range masterMenu.Menus {
			menuParams := saveMenuParams{
				masterMenu: &updateMasterMenu,
				menu:       v,
				parentID:   nil,
				order:      i,
			}

			if v.ID > 0 {
				menuParams.ID = &v.ID
			}

			menu, _ := saveMenu(&r, menuParams)

			if len(v.Children) > 0 {
				for k, child := range v.Children {
					val := models.Menu{
						ParentID:     &menu.ID,
						MasterMenuID: updateMasterMenu.ID,
						MenuItemID:   child.MenuItemID,
						Order:        k,
					}

					childParams := saveMenuParams{
						masterMenu: &updateMasterMenu,
						menu:       val,
						parentID:   &menu.ID,
						order:      k,
					}

					if child.ID > 0 {
						childParams.ID = &child.ID
					}

					childMenu, _ := saveMenu(&r, childParams)

					if len(child.Children) > 0 {
						for _, c := range child.Children {
							saveChildrenMenu(&r, saveChildrenMenuParams{
								masterMenu: &updateMasterMenu,
								menu:       c,
								parentID:   &childMenu.ID,
								menus:      c.Children,
							})
						}
					}
				}
			}
		}
	}

	return updateMasterMenu, nil
}

func (r MasterMenuRepository) Delete(id uint) error {
	var masterMenu = models.MasterMenu{}

	err := r.Model(masterMenu).Where("id = ?", id).First(&masterMenu).Delete(&masterMenu).Error

	if err != nil {
		return err
	}

	if err := r.DB.Model(&models.Menu{}).Delete(&models.Menu{}, "master_menu_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

type saveMenuParams struct {
	ID         *uint
	masterMenu *models.MasterMenu
	menu       models.Menu
	parentID   *uint
	order      int
}

type saveChildrenMenuParams struct {
	masterMenu *models.MasterMenu
	menu       models.Menu
	parentID   *uint
	menus      []models.Menu
}

func saveChildrenMenu(r *MasterMenuRepository, params saveChildrenMenuParams) []models.Menu {
	var result = make([]models.Menu, 0)

	saveParams := saveMenuParams{
		masterMenu: params.masterMenu,
		menu:       params.menu,
		parentID:   params.parentID,
		order:      0,
	}

	if params.menu.ID > 0 {
		saveParams.ID = &params.menu.ID
	}

	data, _ := saveMenu(r, saveParams)

	for i, v := range params.menus {
		childParams := saveMenuParams{
			masterMenu: params.masterMenu,
			menu:       v,
			parentID:   &data.ID,
			order:      i,
		}

		if v.ID > 0 {
			childParams.ID = &v.ID
		}

		child, _ := saveMenu(r, childParams)

		if len(v.Children) > 0 {
			saveChildrenMenu(r, saveChildrenMenuParams{
				masterMenu: params.masterMenu,
				menu:       child,
				parentID:   child.ParentID,
				menus:      v.Children,
			})
		}
	}

	result = append(result, data)

	return result
}

func saveMenu(r *MasterMenuRepository, params saveMenuParams) (models.Menu, error) {
	if params.menu.MenuItemID != 0 {
		data := models.Menu{
			ParentID:     params.parentID,
			MasterMenuID: params.masterMenu.ID,
			MenuItemID:   params.menu.MenuItemID,
			Order:        params.menu.Order,
		}

		if params.ID != nil {
			data.ID = *params.ID
		}

		err := r.Model(&data).Save(&data).Error
		if err != nil {
			return models.Menu{}, err
		}

		return data, nil
	}

	return models.Menu{}, nil
}

func normalizeMenus(menu *models.Menu, menus []models.Menu) []models.Menu {
	var result = make([]models.Menu, 0)

	if len(menus) > 0 {
		for _, v := range menus {
			if v.ParentID != nil {
				if *v.ParentID == menu.ID {
					var children = make([]models.Menu, 0)
					childMenus := normalizeMenus(&v, menus)
					children = append(children, childMenus...)

					v.Children = children
					result = append(result, v)
				}
			}
		}
	}

	return result
}
