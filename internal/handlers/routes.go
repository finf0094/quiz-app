package handlers

import (
	"net/http"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/quizzes", GetQuizzes)
    // r.POST("/responses", CreateResponse)
}


func GetQuizzes(c *gin.Context) {
    // Заглушка
    c.JSON(http.StatusOK, gin.H{
        "message": "List of quizzes",
    })
}
