package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func ValidateHash(hash string, password string) (bool, error) {
	temp, _ := GenerateHash(password)
	err := bcrypt.CompareHashAndPassword([]byte(temp), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
