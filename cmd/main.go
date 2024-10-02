package main

import (
	"fmt"
	_ "github.com/finf0094/quiz-app/docs"
	"github.com/finf0094/quiz-app/internal/config"
	"github.com/finf0094/quiz-app/internal/db"
	"github.com/finf0094/quiz-app/internal/handlers"
	"github.com/finf0094/quiz-app/internal/repositories"
	"github.com/finf0094/quiz-app/internal/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
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

	// Инициализация репозиториев и сервисов
	quizRepo := repositories.NewQuizRepository(db.DB)
	responseRepo := repositories.NewResponseRepository(db.DB)

	quizService := services.NewQuizService(quizRepo)
	responseService := services.NewResponseService(responseRepo)

	// Настройка маршрутов
	handlers.SetupRoutes(r, quizService, responseService)

	// Добавление Swagger маршрута
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Доверительные прокси сервер например локальный прокси (nginx, apache, cloudflare)
	r.SetTrustedProxies([]string{"127.0.0.1"})

	// Запуск сервера
	port := config.AppConfig.Server.Port
	log.Printf("Starting server on port %d...\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
