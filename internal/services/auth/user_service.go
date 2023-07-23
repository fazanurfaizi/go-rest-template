package authServices

import (
	baseModel "github.com/fazanurfaizi/go-rest-template/internal/models"
	authModels "github.com/fazanurfaizi/go-rest-template/internal/models/auth"
	authRepositories "github.com/fazanurfaizi/go-rest-template/internal/repositories/auth"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"gorm.io/gorm"
)

type UserService struct {
	logger          logger.Logger
	repository      authRepositories.UserRepository
	paginationScope *gorm.DB
}

func NewUserService(logger logger.Logger, repository authRepositories.UserRepository) *UserService {
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

func (s UserService) FindById(id baseModel.BinaryUUID) (user authModels.User, err error) {
	return user, s.repository.First(&user, "id = ? ", id).Error
}

func (s UserService) FindAll() (response []authModels.User, total int64, err error) {
	var users []authModels.User
	var count int64

	err = s.repository.WithTrx(s.paginationScope).Find(&users).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (s UserService) Create(user *authModels.User) error {
	return s.repository.Create(&user).Error
}

func (s UserService) Update(user *authModels.User) error {
	return s.repository.Save(&user).Error
}

func (s UserService) Delete(id baseModel.BinaryUUID) error {
	return s.repository.Where("id = ?", id).Delete(&authModels.User{}).Error
}
