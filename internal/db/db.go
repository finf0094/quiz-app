package db

import (
    "fmt"
    "log"

    "github.com/finf0094/quiz-app/internal/config"
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
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }
}
