package jwt

import (
	"testing"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/stretchr/testify/assert"
)

const (
	jwtSecret  string = "JWTSECRET"
	jwtIssuer  string = "go-rest"
	jwtExpired int    = 5
)

func TestGenerateToken(t *testing.T) {
	user := models.User{
		Name:     "tester",
		Email:    "tester@mail.com",
		Password: "password",
	}

	jwtService := NewJWTService(jwtSecret, jwtIssuer, jwtExpired)
	token, err := jwtService.GenerateToken(&user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
