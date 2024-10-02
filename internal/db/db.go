package db

import (
	"fmt"
	"log"

	"github.com/finf0094/quiz-app/internal/config"
	"github.com/finf0094/quiz-app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.AppConfig.Database.Host,
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Name,
		config.AppConfig.Database.Port,
	)

	log.Println("Connecting to the database...")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Successfully connected to the database.")
}

func Migrate() {
	log.Println("Starting database migrations...")
	err := DB.AutoMigrate(
		&models.Quiz{},
		&models.Question{},
		&models.Option{},
		&models.Response{},
		&models.Answer{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Database migrations completed successfully.")
}
