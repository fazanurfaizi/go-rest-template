package postgres

import (
	"fmt"
	"log"

	"github.com/fazanurfaizi/go-rest-template/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewPosgresDB(config *config.Config) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", config.Postgres.PostgresqlHost, config.Postgres.PostgresqlUser, config.Postgres.PostgresqlPassword, config.Postgres.PostgresqlDbName, config.Postgres.PostgresqlPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}

	fmt.Println("Connected successfully to database")
}
