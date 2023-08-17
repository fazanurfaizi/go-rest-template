package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres/filter"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository struct {
	postgres.Database
	logger logger.Logger
}

func NewUserRepository(db postgres.Database, logger logger.Logger) UserRepository {
	return UserRepository{db, logger}
}

func (u UserRepository) WithTrx(trx *gorm.DB) UserRepository {
	if trx != nil {
		u.logger.Debug("Using WithTrx as trxHandle is not nil")
		u.Database.DB = trx
	}
	return u
}

func (u UserRepository) FindAll(ctx *gin.Context) ([]models.User, int64, error) {
	var users []models.User
	var count int64

	err := u.Model(&models.User{}).
		Scopes(filter.FilterByQuery(ctx, filter.ALL)).
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error

	if err != nil {
		return users, 0, err
	}

	return users, count, nil
}
