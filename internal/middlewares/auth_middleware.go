package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/errors"
	"github.com/fazanurfaizi/go-rest-template/pkg/jwt"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	config *config.Config
	logger logger.Logger
}

func NewAuthMiddleware(config *config.Config, logger logger.Logger) *AuthMiddleware {
	return &AuthMiddleware{config: config, logger: logger}
}

func (m AuthMiddleware) Handle() gin.HandlerFunc {
	privateKey, err := os.ReadFile("ssl/id_rsa")
	if err != nil {
		log.Fatalln(err)
	}

	publicKey, err := os.ReadFile("ssl/id_rsa.pub")
	if err != nil {
		log.Fatalln(err)
	}

	jwtService := jwt.NewJWTService(m.config, privateKey, publicKey)

	return func(ctx *gin.Context) {
		claims, err := jwtService.ExtractJWTFromRequest(ctx.Request)
		if err != nil {
			utils.ErrorJSON(ctx, http.StatusUnauthorized, errors.ErrUnauthorized.Error())
			ctx.Abort()
			return
		}

		ctx.Set("user", claims["id"])
		ctx.Next()
	}
}
