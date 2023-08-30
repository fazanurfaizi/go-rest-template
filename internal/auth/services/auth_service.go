package services

import (
	"log"
	"os"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/jwt"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type AuthService interface {
	Login(dto.LoginRequest) (dto.LoginResponse, errors.RestErr)
	Register(dto.RegisterRequesst) (dto.RegisterResponse, errors.RestErr)
}

type authService struct {
	config         *config.Config
	logger         logger.Logger
	userRepository repositories.UserRepository
	jwtService     jwt.JWTService
}

func NewAuthService(
	config *config.Config,
	logger logger.Logger,
	userRepository repositories.UserRepository,
) AuthService {
	privateKey, err := os.ReadFile("ssl/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}

	publicKey, err := os.ReadFile("ssl/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	jwtService := jwt.NewJWTService(config, privateKey, publicKey)

	return &authService{
		config:         config,
		logger:         logger,
		userRepository: userRepository,
		jwtService:     jwtService,
	}
}

func (s *authService) Login(request dto.LoginRequest) (dto.LoginResponse, errors.RestErr) {
	var response = dto.LoginResponse{}
	var user = models.User{}

	user, err := s.userRepository.FindByEmail(request.Email)
	if err != nil {
		return response, errors.NewNotFoundError(err.Error())
	}

	validated, err := user.ComparePassword(request.Password)
	if err != nil {
		return response, errors.NewNotFoundError(err.Error())
	}

	if validated {
		response.User = user
		response.Token, err = s.jwtService.GenerateToken(&jwt.JWTDto{
			ID:    user.ID,
			Email: user.Email,
		})

		if err != nil {
			return response, errors.NewInternalServerError(err.Error())
		}

		return response, nil
	}

	return response, errors.NewInternalServerError(errors.ErrInvalidJWTToken)
}

func (s *authService) Register(request dto.RegisterRequesst) (dto.RegisterResponse, errors.RestErr) {
	panic("Method not implemented!")
}
