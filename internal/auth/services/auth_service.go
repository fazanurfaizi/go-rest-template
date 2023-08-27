package services

import (
	"fmt"
	"log"
	"os"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/dto"
	"github.com/fazanurfaizi/go-rest-template/internal/auth/repositories"
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/jwt"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type AuthService struct {
	config     *config.Config
	logger     logger.Logger
	repository repositories.AuthRepository
	jwtService jwt.JWTService
}

func NewAuthService(
	config *config.Config,
	logger logger.Logger,
	repository repositories.AuthRepository,
) *AuthService {
	privateKey, err := os.ReadFile("ssl/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(privateKey)

	publicKey, err := os.ReadFile("ssl/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	jwtService := jwt.NewJWTService(privateKey, publicKey, config.Server.AppName, 10)

	return &AuthService{
		config:     config,
		logger:     logger,
		repository: repository,
		jwtService: jwtService,
	}
}

func (s AuthService) Login(request dto.LoginRequest) (dto.LoginResponse, errors.RestErr) {
	user, err := s.repository.FindByEmail(request.Email)
	if err != nil {
		return dto.LoginResponse{}, errors.NewNotFoundError(err.Error())
	}

	validated, err := user.ComparePassword(request.Password)
	if err != nil {
		return dto.LoginResponse{}, errors.NewNotFoundError(err.Error())
	}

	if validated {
		token, err := s.jwtService.GenerateToken(&jwt.JWTDto{
			ID:    user.ID,
			Email: user.Email,
		})

		if err != nil {
			return dto.LoginResponse{}, errors.NewInternalServerError(err.Error())
		}

		return dto.LoginResponse{
			User:  user,
			Token: token,
		}, nil
	}

	return dto.LoginResponse{}, errors.NewInternalServerError(errors.ErrInvalidJWTToken)
}
