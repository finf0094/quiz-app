package repositories

import (
	"github.com/finf0094/quiz-app/internal/models"
	"gorm.io/gorm"
)

// QuizRepository - интерфейс для работы с викторинами
type QuizRepository interface {
	CreateQuiz(quiz *models.Quiz) error
	GetQuizByID(id uint) (*models.Quiz, error)
	GetAllQuizzes() ([]models.Quiz, error)
	UpdateQuiz(quiz *models.Quiz) error
	DeleteQuiz(id uint) error
}

// NewQuizRepository - создание нового экземпляра репозитория викторин
func NewQuizRepository(db *gorm.DB) QuizRepository {
	return &GormQuizRepository{db: db}
}

// GormQuizRepository - реализация QuizRepository с использованием Gorm
type GormQuizRepository struct {
	db *gorm.DB
}

func (r *GormQuizRepository) CreateQuiz(quiz *models.Quiz) error {
	return r.db.Create(quiz).Error
}

func (r *GormQuizRepository) GetQuizByID(id uint) (*models.Quiz, error) {
	var quiz models.Quiz
	err := r.db.Preload("Questions.Options").First(&quiz, id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *GormQuizRepository) GetAllQuizzes() ([]models.Quiz, error) {
	var quizzes []models.Quiz
	err := r.db.Preload("Questions.Options").Find(&quizzes).Error
	if err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (r *GormQuizRepository) UpdateQuiz(quiz *models.Quiz) error {
	return r.db.Save(quiz).Error
}

func (r *GormQuizRepository) DeleteQuiz(id uint) error {
	return r.db.Delete(&models.Quiz{}, id).Error
}
