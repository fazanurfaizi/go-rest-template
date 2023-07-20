package jwt

import (
	"errors"
	"time"

	jwtDriver "github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId string, username string) (token string, err error)
	ParseToken(token string) (claims JwtCustomClaim, err error)
}

type JwtCustomClaim struct {
	UserId   string
	Username string
	jwtDriver.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
	expired   int
}

func NewJWTService(secretKey string, issuer string, expired int) JWTService {
	return &jwtService{
		secretKey: secretKey,
		issuer:    issuer,
		expired:   expired,
	}
}

func (j *jwtService) GenerateToken(userId string, username string) (token string, err error) {
	claims := &JwtCustomClaim{
		userId,
		username,
		jwtDriver.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.expired)).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwtDriver.NewWithClaims(jwtDriver.SigningMethodHS256, claims)
	t, err := jwtToken.SignedString([]byte(j.secretKey))
	return t, err
}

func (j *jwtService) ParseToken(token string) (claims JwtCustomClaim, err error) {
	if jwtToken, err := jwtDriver.ParseWithClaims(token, &claims, func(t *jwtDriver.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	}); err != nil || !jwtToken.Valid {
		return JwtCustomClaim{}, errors.New("token is not valid")
	}

	return
}
