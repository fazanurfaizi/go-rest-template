package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"gorm.io/gorm/clause"
)

type AuthRepository struct {
	postgres.Database
	logger logger.Logger
}

func NewAuthRepository(db postgres.Database, logger logger.Logger) AuthRepository {
	return AuthRepository{db, logger}
}

func (r *AuthRepository) Register(user *models.User) (models.User, error) {
	err := r.Model(user).Save(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return *user, nil
}

func (r *AuthRepository) Update(id uint, user *models.User) (models.User, error) {
	var updatedUser = models.User{}
	err := r.Model(&updatedUser).Clauses(clause.Returning{}).Where("id = ?", id).First(&updatedUser).Updates(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return updatedUser, nil
}

func (r *AuthRepository) FindById(id uint) (models.User, error) {
	var user = models.User{}
	err := r.Model(user).Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *AuthRepository) FindByEmail(email string) (models.User, error) {
	var user = models.User{}
	err := r.Model(user).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
