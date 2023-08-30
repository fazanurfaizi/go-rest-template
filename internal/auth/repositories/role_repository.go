package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres/filter"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type RoleRepository interface {
	FindAll(*gin.Context) ([]models.Role, int64)
	FindById(uint) (models.Role, error)
	Create(*models.Role) (models.Role, error)
	Update(uint, *models.Role) (models.Role, error)
	Delete(uint) error
}

type roleRepository struct {
	postgres.Database
	logger logger.Logger
}

func NewRoleRepository(db postgres.Database, logger logger.Logger) RoleRepository {
	return &roleRepository{db, logger}
}

func (r *roleRepository) FindAll(ctx *gin.Context) ([]models.Role, int64) {
	var roles []models.Role
	var count int64

	r.Model(&models.Role{}).
		Scopes(filter.FilterByQuery(ctx, filter.ALL)).
		Find(&roles).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return roles, count
}

func (r *roleRepository) FindById(id uint) (role models.Role, err error) {
	err = r.Model(&role).Preload("Permissions").Where("id = ?", id).First(&role).Error
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func (r *roleRepository) Create(role *models.Role) (models.Role, error) {
	err := r.Model(role).Save(&role).Error
	if err != nil {
		return models.Role{}, err
	}

	return *role, nil
}

func (r *roleRepository) Update(id uint, role *models.Role) (models.Role, error) {
	var updatedRole = models.Role{}
	err := r.Model(&updatedRole).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		First(&role).
		Updates(&role).
		Error

	if err != nil {
		return models.Role{}, err
	}

	if err := r.Model(&role).Association("Permissions").Replace(role.Permissions); err != nil {
		return models.Role{}, err
	}

	return updatedRole, nil
}

func (r *roleRepository) Delete(id uint) error {
	var role = models.Role{}

	err := r.Model(role).Where("id = ?", id).First(&role).Delete(&role).Error

	if err != nil {
		return err
	}

	// if err := r.Model(&role).Association("Permissions").Clear(); err != nil {
	// 	return err
	// }

	return nil
}
