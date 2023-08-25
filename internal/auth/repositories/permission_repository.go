package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres/filter"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type PermissionRepository struct {
	postgres.Database
	logger logger.Logger
}

func NewPermissionRepository(db postgres.Database, logger logger.Logger) PermissionRepository {
	return PermissionRepository{db, logger}
}

func (r PermissionRepository) FindAll(ctx *gin.Context) ([]models.Permission, int64) {
	var permissions []models.Permission
	var count int64

	r.Model(&models.Permission{}).
		Scopes(filter.FilterByQuery(ctx, filter.ALL)).
		Find(&permissions).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return permissions, count
}

func (r PermissionRepository) FindById(id uint) (permission models.Permission, err error) {
	err = r.Model(permission).Where("id = ?", id).First(&permission).Error
	if err != nil {
		return models.Permission{}, err
	}
	return permission, nil
}

func (r PermissionRepository) Create(permission *models.Permission) (models.Permission, error) {
	err := r.Model(permission).Save(&permission).Error
	if err != nil {
		return models.Permission{}, err
	}

	return *permission, nil
}

func (r PermissionRepository) Update(id uint, permission *models.Permission) (models.Permission, error) {
	var updatedPermission = models.Permission{}
	err := r.Model(&updatedPermission).Clauses(clause.Returning{}).Where("id = ?", id).First(&permission).Updates(&permission).Error
	if err != nil {
		return models.Permission{}, err
	}

	return updatedPermission, nil
}

func (r PermissionRepository) Delete(id uint) error {
	var permission = models.Permission{}
	err := r.Model(permission).Where("id = ?", id).First(&permission).Delete(&permission).Error
	if err != nil {
		return err
	}

	return nil
}
