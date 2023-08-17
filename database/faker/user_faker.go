package faker

import (
	"github.com/brianvoe/gofakeit/v6"
	authModels "github.com/fazanurfaizi/go-rest-template/internal/models/auth"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

type UserFaker struct {
	logger logger.Logger
	db     postgres.Database
}

func NewUserFaker(logger logger.Logger, db postgres.Database) UserFaker {
	return UserFaker{
		logger: logger,
		db:     db,
	}
}

func (f UserFaker) Setup() {
	f.logger.Info("User faker data...")

	// var users []authModels.User
	// _ = faker.FakeData(&sample)
	var user authModels.User
	gofakeit.Struct(&user)
	// for i := 0; i < 100; i++ {
	// }

	f.logger.Info("User faker data successfully")
}
