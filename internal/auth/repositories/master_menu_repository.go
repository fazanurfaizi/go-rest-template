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
	err = r.Model(&masterMenu).Preload("Menus").Where("id = ?", id).First(&masterMenu).Error
	if err != nil {
		return models.MasterMenu{}, err
	}

	return masterMenu, nil
}

func (r MasterMenuRepository) Create(masterMenu *models.MasterMenu) (models.MasterMenu, error) {
	createdMasterMenu := models.MasterMenu{Name: masterMenu.Name}

	result := r.Model(&createdMasterMenu).Clauses(clause.Returning{}).Select("id", "name").Create(&createdMasterMenu)

	if result.Error != nil {
		return models.MasterMenu{}, result.Error
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
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		First(&masterMenu).
		Updates(&masterMenu).
		Error

	if err != nil {
		return models.MasterMenu{}, err
	}

	// if err := r.Model(&masterMenu).Association("Menus").Replace(masterMenu.Menus); err != nil {
	// 	return models.MasterMenu{}, err
	// }

	return updateMasterMenu, nil
}

func (r MasterMenuRepository) Delete(id uint) error {
	var masterMenu = models.MasterMenu{}

	err := r.Model(masterMenu).Where("id = ?", id).First(&masterMenu).Delete(&masterMenu).Error

	if err != nil {
		return err
	}

	// if err := r.Model(&masterMenu).Association("Permissions").Clear(); err != nil {
	// 	return err
	// }

	return nil
}

type saveMenuParams struct {
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

	data, _ := saveMenu(r, saveMenuParams{
		masterMenu: params.masterMenu,
		menu:       params.menu,
		parentID:   params.parentID,
		order:      0,
	})

	for i, v := range params.menus {
		child, _ := saveMenu(r, saveMenuParams{
			masterMenu: params.masterMenu,
			menu:       v,
			parentID:   &data.ID,
			order:      i,
		})

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

		err := r.Model(&data).Create(&data).Error
		if err != nil {
			return models.Menu{}, err
		}

		return data, nil
	}

	return models.Menu{}, nil
}
