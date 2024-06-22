package main

import (
	"fmt"
	"log"

	"github.com/khusanov-m/qolmaqol-api/initializers"
	"github.com/khusanov-m/qolmaqol-api/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		panic("failed to auto-migrate")
	}
	fmt.Println("? Migration complete")
}
