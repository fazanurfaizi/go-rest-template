package jwt

import (
	"testing"

	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	baseModel "github.com/fazanurfaizi/go-rest-template/internal/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	jwtSecret  string = "JWTSECRET"
	jwtIssuer  string = "go-rest"
	jwtExpired int    = 5
)

func TestGenerateToken(t *testing.T) {
	id, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}

	user := models.User{
		Name:     "tester",
		Email:    "tester@mail.com",
		Password: "password",
	}
	user.BaseModel.ID = baseModel.BinaryUUID(id)

	jwtService := NewJWTService(jwtSecret, jwtIssuer, jwtExpired)
	token, err := jwtService.GenerateToken(&user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
