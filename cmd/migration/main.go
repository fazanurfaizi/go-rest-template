package main

import (
	"log"
)

func main() {
	// log.Println("Load config files")
	// configPath := utils.GetConfigPath(os.Getenv("config"))

	// configFile, err := config.LoadConfig(configPath)
	// if err != nil {
	// 	log.Fatalf("LoadConfig: %v", err)
	// }

	// config, err := config.ParseConfig(configFile)
	// if err != nil {
	// 	log.Fatalf("ParsConfig: %v", err)
	// }

	// database := postgres.NewConnection(config)

	// schemas := [3]string{"auth", "master", "transaction"}
	// for _, schema := range schemas {
	// 	database.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", schema))
	// }

	// database.AutoMigrate(&models.User{})
	log.Println("Migration complete")
}
