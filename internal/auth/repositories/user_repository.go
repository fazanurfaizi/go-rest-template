package repositories

import (
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
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
