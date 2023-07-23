package postgres

import (
	"fmt"
	"log"

	"github.com/fazanurfaizi/go-rest-template/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewConnection(config *config.Config) Database {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Postgres.PostgresqlHost, config.Postgres.PostgresqlUser, config.Postgres.PostgresqlPassword, config.Postgres.PostgresqlDbName, config.Postgres.PostgresqlPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}

	fmt.Println("Connected successfully to database")

	return Database{DB: db}
}
