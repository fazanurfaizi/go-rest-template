package jwt

import (
	"errors"
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
)

type JWTService interface {
	GenerateToken(user *JWTDto) (token string, err error)
	ExtractJWTFromRequest(r *http.Request) (map[string]interface{}, error)
}

type Claims struct {
	Email string `json:"email"`
	ID    uint   `json:"id"`
	jwt.StandardClaims
}

type jwtService struct {
	config     *config.Config
	privateKey []byte
	publicKey  []byte
}

type JWTDto struct {
	ID    uint
	Email string
}

func NewJWTService(
	config *config.Config,
	privateKey []byte,
	publicKey []byte,
) JWTService {
	return &jwtService{
		config:     config,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (j *jwtService) GenerateToken(user *JWTDto) (token string, err error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", errors.New(err.Error())
	}

	claims := &Claims{
		Email: user.Email,
		ID:    user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.config.Session.Expire)).Unix(),
			Issuer:    j.config.Session.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	t, err := jwtToken.SignedString(key)
	return t, err
}

// Extract JWT from request
func (j *jwtService) ExtractJWTFromRequest(r *http.Request) (map[string]interface{}, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.publicKey)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	// Get the jwt string
	tokenString := ExtractBearerToken(r)

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(err.Error())
		}

		return key, nil
	})

	if err != nil {
		return nil, errors.New(err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New(err.Error())
	}

	return claims, nil
}

// Extract bearer token from request Authorization header
func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	if len(bearerToken) > 1 {
		return html.EscapeString(bearerToken[1])
	}
	return ""
}
