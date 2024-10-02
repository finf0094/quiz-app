package handlers

import (
	"net/http"

	"github.com/finf0094/quiz-app/internal/models"
	"github.com/finf0094/quiz-app/internal/services"
	"github.com/gin-gonic/gin"
)

type ResponseHandler struct {
	responseService services.ResponseService
}

func NewResponseHandler(responseService services.ResponseService) *ResponseHandler {
	return &ResponseHandler{responseService: responseService}
}

// CreateResponse создает новый ответ пользователя на викторину
// @Summary Create a new response
// @Description Create a new response from a user for a quiz
// @Tags responses
// @Accept json
// @Produce json
// @Param response body models.Response true "Response to create"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /responses [post]
func (h *ResponseHandler) CreateResponse(c *gin.Context) {
	var response models.Response
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.responseService.CreateResponse(&response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}
