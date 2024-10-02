package services

import (
	"github.com/finf0094/quiz-app/internal/models"
	"github.com/finf0094/quiz-app/internal/repositories"
)

type QuizService interface {
	CreateQuiz(quiz *models.Quiz) error
	GetQuizByID(id uint) (*models.Quiz, error)
	GetAllQuizzes() ([]models.Quiz, error)
	UpdateQuiz(quiz *models.Quiz) error
	DeleteQuiz(id uint) error
}

type quizService struct {
	quizRepo repositories.QuizRepository
}

func NewQuizService(quizRepo repositories.QuizRepository) QuizService {
	return &quizService{quizRepo: quizRepo}
}

func (s *quizService) CreateQuiz(quiz *models.Quiz) error {
	return s.quizRepo.CreateQuiz(quiz)
}

func (s *quizService) GetQuizByID(id uint) (*models.Quiz, error) {
	return s.quizRepo.GetQuizByID(id)
}

func (s *quizService) GetAllQuizzes() ([]models.Quiz, error) {
	return s.quizRepo.GetAllQuizzes()
}

func (s *quizService) UpdateQuiz(quiz *models.Quiz) error {
	return s.quizRepo.UpdateQuiz(quiz)
}

func (s *quizService) DeleteQuiz(id uint) error {
	return s.quizRepo.DeleteQuiz(id)
}
