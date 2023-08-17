package services

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	logger          logger.Logger
	repository      repositories.UserRepository
	paginationScope *gorm.DB
}

func NewUserService(logger logger.Logger, repository repositories.UserRepository) *UserService {
	return &UserService{
		logger:     logger,
		repository: repository,
	}
}

func (s UserService) WithTrx(trx *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trx)
	return s
}

func (s UserService) SetPaginationScope(scope func(*gorm.DB) *gorm.DB) UserService {
	s.paginationScope = s.repository.WithTrx(s.repository.Scopes(scope)).DB
	return s
}

func (s UserService) FindById(id uint) models.User {
	return s.repository.FindById(id)
}

func (s UserService) FindByEmailAndPassword(email string, password string) (user models.User, err error) {
	s.repository.First(&user, "email = ? ", email)
	_, err = utils.ValidateHash(password, user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s UserService) FindAll(ctx *gin.Context) ([]models.User, int64) {
	return s.repository.FindAll(ctx)
}

func (s UserService) Create(user *models.User) error {
	return s.repository.Create(&user).Error
}

func (s UserService) Update(user *models.User) error {
	return s.repository.Save(&user).Error
}

func (s UserService) Delete(id uint) error {
	return s.repository.Where("id = ?", id).Delete(&models.User{}).Error
}
