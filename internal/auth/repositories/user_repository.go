package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres/filter"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	FindAll(*gin.Context) ([]models.User, int64)
	FindById(uint) (models.User, error)
	FindByEmail(string) (models.User, error)
	Create(*models.User) (models.User, error)
	Update(uint, *models.User) (models.User, error)
	Delete(uint) error
}

type userRepository struct {
	postgres.Database
	logger logger.Logger
}

func NewUserRepository(db postgres.Database, logger logger.Logger) UserRepository {
	return &userRepository{db, logger}
}

func (u *userRepository) FindAll(ctx *gin.Context) ([]models.User, int64) {
	var users []models.User
	var count int64

	u.Model(&models.User{}).
		Scopes(filter.FilterByQuery(ctx, filter.ALL)).
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&count)

	return users, count
}

func (u *userRepository) FindById(id uint) (models.User, error) {
	var user = models.User{}
	err := u.Model(user).Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) FindByEmail(email string) (models.User, error) {
	var user = models.User{}
	err := u.Model(user).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *userRepository) Create(user *models.User) (models.User, error) {
	err := u.Model(user).Save(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (u *userRepository) Update(id uint, user *models.User) (models.User, error) {
	var updatedUser = models.User{}
	err := u.Model(&updatedUser).Clauses(clause.Returning{}).Where("id = ?", id).First(&updatedUser).Updates(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (u *userRepository) Delete(id uint) error {
	var user = models.User{}
	err := u.Model(user).Where("id = ?", id).First(&user).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
