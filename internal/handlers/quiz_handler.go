package handlers

import (
	"net/http"
	"strconv"

	"github.com/finf0094/quiz-app/internal/models"
	"github.com/finf0094/quiz-app/internal/services"
	"github.com/gin-gonic/gin"
)

type QuizHandler struct {
	quizService services.QuizService
}

func NewQuizHandler(quizService services.QuizService) *QuizHandler {
	return &QuizHandler{quizService: quizService}
}

// GetQuizzes получает все викторины
// @Summary Get all quizzes
// @Description Get a list of all quizzes
// @Tags quizzes
// @Produce json
// @Success 200 {array} models.Quiz
// @Failure 500 {object} models.ErrorResponse
// @Router /quizzes [get]
func (h *QuizHandler) GetQuizzes(c *gin.Context) {
	quizzes, err := h.quizService.GetAllQuizzes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

// GetQuizByID получает викторину по ID
// @Summary Get a quiz by ID
// @Description Get a quiz by its unique ID
// @Tags quizzes
// @Param id path int true "Quiz ID"
// @Produce json
// @Success 200 {object} models.Quiz
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /quizzes/{id} [get]
func (h *QuizHandler) GetQuizByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	quiz, err := h.quizService.GetQuizByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

// CreateQuiz создает новую викторину
// @Summary Create a new quiz
// @Description Create a new quiz with the given details
// @Tags quizzes
// @Accept json
// @Produce json
// @Param quiz body models.Quiz true "Quiz to create"
// @Success 201 {object} models.Quiz
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /quizzes [post]
func (h *QuizHandler) CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.quizService.CreateQuiz(&quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, quiz)
}
