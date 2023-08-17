package services

import (
	"mime/multipart"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/storage"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/formatter"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	logger          logger.Logger
	repository      repositories.UserRepository
	fileStorage     storage.FileStorage
	paginationScope *gorm.DB
}

func NewUserService(logger logger.Logger, repository repositories.UserRepository, fileStorage storage.FileStorage) *UserService {
	return &UserService{
		logger:      logger,
		repository:  repository,
		fileStorage: fileStorage,
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

func (s UserService) FindById(id uint) (dto.UserResponse, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return dto.UserResponse{}, err
	}

	avatarUrl, _ := s.fileStorage.GetFile(user.Avatar)
	return dto.MappingUserResponse(user, avatarUrl), nil
}

func (s UserService) FindByEmailAndPassword(email string, password string) (user models.User, err error) {
	s.repository.First(&user, "email = ? ", email)
	_, err = utils.ValidateHash(password, user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s UserService) FindAll(ctx *gin.Context) ([]dto.UserResponse, int64) {
	var result []dto.UserResponse
	users, total := s.repository.FindAll(ctx)
	for _, user := range users {
		avatarUrl, _ := s.fileStorage.GetFile(user.Avatar)
		result = append(result, dto.MappingUserResponse(user, avatarUrl))
	}

	return result, total
}

func (s UserService) Create(request dto.CreateUserRequest, file multipart.File) (dto.UserResponse, errors.RestErr) {
	filename, errStorage := s.fileStorage.Upload(file)
	if errStorage != nil {
		return dto.UserResponse{}, errors.NewBadRequestError(errStorage.Error())
	}

	user := models.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
		Avatar:      filename,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		City:        request.City,
		Country:     request.Country,
		Gender:      request.Gender,
		Postcode:    request.Postcode,
		Birthday:    formatter.ParseStringToTime(request.Birthday, formatter.YYYYMMDDhhmmss),
	}
	err := user.HashPassword()
	if err != nil {
		return dto.UserResponse{}, errors.NewInternalServerError("error while hashing password")
	}

	err = s.repository.Create(&user)
	if err != nil {
		return dto.UserResponse{}, errors.NewBadRequestError(err.Error())
	}

	avatarUrl, _ := s.fileStorage.GetFile(user.Avatar)

	return dto.MappingUserResponse(user, avatarUrl), nil
}

func (s UserService) Update(user *models.User) error {
	return s.repository.Save(&user).Error
}

func (s UserService) Delete(id uint) error {
	return s.repository.Where("id = ?", id).Delete(&models.User{}).Error
}
