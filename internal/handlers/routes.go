package handlers

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/quizzes", GetQuizzes)
    // r.POST("/responses", CreateResponse)
    // Добавьте другие маршруты для вашего приложения
}


func GetQuizzes(c *gin.Context) {
    // Заглушка - здесь будет обращение к базе данных через сервис
    c.JSON(http.StatusOK, gin.H{
        "message": "List of quizzes",
    })
}
