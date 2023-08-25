package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres/filter"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type MenuItemRepository struct {
	postgres.Database
	logger logger.Logger
}

func NewMenuItemRepository(db postgres.Database, logger logger.Logger) MenuItemRepository {
	return MenuItemRepository{db, logger}
}

func (r MenuItemRepository) FindAll(ctx *gin.Context) ([]models.MenuItem, int64) {
	var menuItems []models.MenuItem
	var count int64

	r.Model(&models.MenuItem{}).
		Scopes(filter.FilterByQuery(ctx, filter.ALL)).
		Find(&menuItems).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return menuItems, count
}

func (r MenuItemRepository) FindById(id uint) (menuItem models.MenuItem, err error) {
	err = r.Model(menuItem).Where("id = ?", id).First(&menuItem).Error
	if err != nil {
		return models.MenuItem{}, err
	}
	return menuItem, nil
}

func (r MenuItemRepository) Create(menuItem *models.MenuItem) (models.MenuItem, error) {
	err := r.Model(menuItem).Save(&menuItem).Error
	if err != nil {
		return models.MenuItem{}, err
	}

	return *menuItem, nil
}

func (r MenuItemRepository) Update(id uint, menuItem *models.MenuItem) (models.MenuItem, error) {
	var updatedMenuItem = models.MenuItem{}
	err := r.Model(&updatedMenuItem).Clauses(clause.Returning{}).Where("id = ?", id).First(&menuItem).Updates(&menuItem).Error
	if err != nil {
		return models.MenuItem{}, err
	}

	return updatedMenuItem, nil
}

func (r MenuItemRepository) Delete(id uint) error {
	var menuItem = models.MenuItem{}
	err := r.Model(menuItem).Where("id = ?", id).First(&menuItem).Delete(&menuItem).Error
	if err != nil {
		return err
	}

	return nil
}
