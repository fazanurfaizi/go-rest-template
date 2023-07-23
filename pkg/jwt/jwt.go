package jwt

import (
	"errors"
	"html"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	authModels "github.com/fazanurfaizi/go-rest-template/internal/models/auth"
)

type JWTService interface {
	GenerateToken(user *authModels.User) (token string, err error)
	ExtractJWTFromRequest(r *http.Request) (map[string]interface{}, error)
}

type Claims struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	jwt.StandardClaims
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

func (j *jwtService) GenerateToken(user *authModels.User) (token string, err error) {
	claims := &Claims{
		Email: user.Email,
		ID:    user.ID.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.expired)).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := jwtToken.SignedString([]byte(j.secretKey))
	return t, err
}

// Extract JWT from request
func (j *jwtService) ExtractJWTFromRequest(r *http.Request) (map[string]interface{}, error) {
	// Get the jwt string
	tokenString := ExtractBearerToken(r)

	// Initialize a new instance of Claims
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (jwtKey interface{}, err error) {
		return jwtKey, err
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("invalid token signatur")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// Extract bearer token from request Authorization header
func ExtractBearerToken(r *http.Request) string {
	headerAuthorization := r.Header.Get("Authorization")
	bearerToken := strings.Split(headerAuthorization, " ")
	return html.EscapeString(bearerToken[1])
}
