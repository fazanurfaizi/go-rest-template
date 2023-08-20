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
	user := models.User{
		Name:        request.Name,
		Email:       request.Email,
		Password:    request.Password,
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

	if file != nil {
		filename, err := s.fileStorage.Upload(file)
		if err != nil {
			return dto.UserResponse{}, errors.NewInternalServerError("error while uploading avatar")
		}
		user.Avatar = filename
	}

	createdUser, err := s.repository.Create(&user)
	if err != nil {
		return dto.UserResponse{}, errors.NewBadRequestError(err.Error())
	}

	avatarUrl, _ := s.fileStorage.GetFile(user.Avatar)

	return dto.MappingUserResponse(createdUser, avatarUrl), nil
}

func (s UserService) Update(id uint, request dto.UpdateUserRequest, file multipart.File) (dto.UserResponse, errors.RestErr) {
	user := models.User{
		Name:        request.Name,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Address:     request.Address,
		City:        request.City,
		Country:     request.Country,
		Gender:      request.Gender,
		Postcode:    request.Postcode,
		Birthday:    formatter.ParseStringToTime(request.Birthday, formatter.YYYYMMDDhhmmss),
	}

	if request.Password != "" {
		user.Password = request.Password
		err := user.HashPassword()
		if err != nil {
			return dto.UserResponse{}, errors.NewInternalServerError("error while hashing password")
		}
	}

	if file != nil {
		filename, err := s.fileStorage.Upload(file)
		if err != nil {
			return dto.UserResponse{}, errors.NewInternalServerError("error while uploading avatar")
		}
		user.Avatar = filename
	}

	updatedUser, err := s.repository.Update(id, &user)
	if err != nil {
		return dto.UserResponse{}, nil
	}

	avatarUrl, _ := s.fileStorage.GetFile(updatedUser.Avatar)

	return dto.MappingUserResponse(updatedUser, avatarUrl), nil
}

func (s UserService) Delete(id uint) errors.RestErr {
	err := s.repository.Delete(id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	return nil
}
