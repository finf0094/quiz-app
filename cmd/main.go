package main

import (
    "fmt"
    "log"
    "github.com/gin-gonic/gin"
    "github.com/finf0094/quiz-app/internal/config"
    "github.com/finf0094/quiz-app/internal/db"
    "github.com/finf0094/quiz-app/internal/handlers"
)

func main() {
    // Загрузка конфигураций
    config.LoadConfig()
    log.Println("Configuration loaded successfully.")

    // Подключение к базе данных
    db.ConnectDatabase()
    log.Println("Connected to the database successfully.")

    // Выполнение миграций
    db.Migrate()

    // Роутер GIN
    r := gin.New()

    // Настройка маршрутов
    handlers.SetupRoutes(r)

    // Доверительные прокси сервер например локальный прокси (nginx, apache, cloudflare)
    r.SetTrustedProxies([]string{"127.0.0.1"})

    // Запуск сервера
    port := config.AppConfig.Server.Port
    log.Printf("Starting server on port %d...\n", port)
    r.Run(fmt.Sprintf(":%d", port))
}
