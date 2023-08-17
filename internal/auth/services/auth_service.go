package services

import "github.com/fazanurfaizi/go-rest-template/pkg/logger"

type AuthService struct {
	logger      logger.Logger
	userService UserService
}

func NewAuthService(logger logger.Logger, userService UserService) *AuthService {
	return &AuthService{logger, userService}
}

func (s AuthService) SignIn() {

}
