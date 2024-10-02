package handlers

import (
	"github.com/finf0094/quiz-app/internal/services"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, quizService services.QuizService, responseService services.ResponseService) {
	quizHandler := NewQuizHandler(quizService)
	responseHandler := NewResponseHandler(responseService)

	// Маршруты для работы с викторинами
	r.GET("/quizzes", quizHandler.GetQuizzes)
	r.GET("/quizzes/:id", quizHandler.GetQuizByID)
	r.POST("/quizzes", quizHandler.CreateQuiz)

	// Маршруты для работы с ответами пользователей
	r.POST("/responses", responseHandler.CreateResponse)
}
