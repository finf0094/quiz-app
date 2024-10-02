package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/finf0094/quiz-app/internal/config"
    "github.com/finf0094/quiz-app/internal/db"
    "github.com/finf0094/quiz-app/internal/handlers"
)

func main() {
    // Загрузите конфигурацию
    config.LoadConfig()

    // Подключитесь к базе данных
    db.ConnectDatabase()

    // Создайте роутер Gin
    r := gin.New()

    // Настройте маршруты
    handlers.SetupRoutes(r)

    // Доверительные прокси сервер например локальный прокси
    r.SetTrustedProxies([]string{"127.0.0.1"})

    // Запустите сервер
    port := config.AppConfig.Server.Port
    r.Run(fmt.Sprintf(":%d", port))
}
