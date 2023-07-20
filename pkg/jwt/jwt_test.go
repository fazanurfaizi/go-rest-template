package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	jwtSecret  string = "JWTSECRET"
	jwtIssuer  string = "go-rest"
	jwtExpired int    = 5
	userId     string = "asf-asf-dsas-sadasdasd"
	username   string = "faza.nurfaizi"
	password   string = "password"
)

func TestGenerateToken(t *testing.T) {
	jwtService := NewJWTService(jwtSecret, jwtIssuer, jwtExpired)
	token, err := jwtService.GenerateToken(userId, username)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	t.Run("With Valid Token", func(t *testing.T) {
		jwtService := NewJWTService(jwtSecret, jwtIssuer, jwtExpired)

		token, err := jwtService.GenerateToken(userId, username)
		assert.NoError(t, err, "Error while generating token")

		claims, err := jwtService.ParseToken(token)
		assert.NoError(t, err, "Errow while parse token")
		assert.Equal(t, userId, claims.UserId)
		assert.Equal(t, username, claims.Username)
		assert.True(t, claims.StandardClaims.ExpiresAt >= time.Now().Unix())
		assert.Equal(t, jwtIssuer, claims.StandardClaims.Issuer)
		assert.True(t, claims.StandardClaims.IssuedAt <= time.Now().Unix())
	})

	t.Run("With Invalid Token", func(t *testing.T) {
		jwtService := NewJWTService(jwtSecret, jwtIssuer, jwtExpired)
		_, err := jwtService.ParseToken("invalid_token")
		assert.Error(t, err)
		assert.Equal(t, "token is not valid", err.Error())
	})
}
