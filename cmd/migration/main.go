package main

import (
	"fmt"
	"log"

	authModels "github.com/fazanurfaizi/go-rest-template/internal/auth/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"github.com/fazanurfaizi/go-rest-template/pkg/core/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/logger"
)

func main() {
	log.Println("Load config files")
	config := config.NewConfig(logger.GetLogger())

	database := postgres.NewConnection(config)

	schemas := [4]string{"public", "auth", "master", "transaction"}
	for _, schema := range schemas {
		database.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schema))
	}

	database.AutoMigrate(
		&authModels.User{},
		&authModels.Permission{},
		&authModels.Role{},
		// &authModels.RolePermission{},
	)
	log.Println("Migration complete")
}
