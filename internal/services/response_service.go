package services

import (
	"github.com/finf0094/quiz-app/internal/models"
	"github.com/finf0094/quiz-app/internal/repositories"
)

type ResponseService interface {
	CreateResponse(response *models.Response) error
	GetResponseByID(id uint) (*models.Response, error)
}

type responseService struct {
	responseRepo repositories.ResponseRepository
}

func NewResponseService(responseRepo repositories.ResponseRepository) ResponseService {
	return &responseService{responseRepo: responseRepo}
}

func (s *responseService) CreateResponse(response *models.Response) error {
	return s.responseRepo.CreateResponse(response)
}

func (s *responseService) GetResponseByID(id uint) (*models.Response, error) {
	return s.responseRepo.GetResponseByID(id)
}
