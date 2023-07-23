package middlewares

import (
	"net/http"

	"github.com/fazanurfaizi/go-rest-template/pkg/constants"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

type DBTransactionMiddleware struct {
	logger logger.Logger
	db     postgres.Database
}

func NewDBTransactionMiddleware(
	logger logger.Logger,
	db postgres.Database,
) DBTransactionMiddleware {
	return DBTransactionMiddleware{
		logger: logger,
		db:     db,
	}
}

func (m DBTransactionMiddleware) Handle() gin.HandlerFunc {
	m.logger.Info("Setting up database transaction middleware")

	return func(ctx *gin.Context) {
		txHandle := m.db.DB.Begin()
		m.logger.Info("Beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		ctx.Set(constants.DBTransaction, txHandle)
		ctx.Next()

		if utils.StatusInList(ctx.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			m.logger.Info("Committing transaction")
			if err := txHandle.Commit().Error; err != nil {
				m.logger.Error("Transaction commit error: ", err)
			}
		} else {
			m.logger.Info("Rolling back transaction due to status code: ", ctx.Writer.Status())
			txHandle.Rollback()
		}
	}
}
