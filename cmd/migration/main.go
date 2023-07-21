package main

import (
	"log"
	"os"

	"github.com/fazanurfaizi/go-rest-template/config"
	"github.com/fazanurfaizi/go-rest-template/internal/models"
	"github.com/fazanurfaizi/go-rest-template/pkg/db/postgres"
	"github.com/fazanurfaizi/go-rest-template/pkg/utils"
)

func init() {
	log.Println("Load config files")
	configPath := utils.GetConfigPath(os.Getenv("config"))

	configFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}

	config, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("ParsConfig: %v", err)
	}

	postgres.NewPosgresDB(config)
}

func main() {
	postgres.DB.AutoMigrate(&models.User{})
	log.Println("Migration complete")
}
