package repositories

import (
	"github.com/finf0094/quiz-app/internal/models"
	"gorm.io/gorm"
)

// ResponseRepository - интерфейс для работы с ответами пользователей
type ResponseRepository interface {
	CreateResponse(response *models.Response) error
	GetResponseByID(id uint) (*models.Response, error)
}

// NewResponseRepository - создание нового экземпляра репозитория ответов
func NewResponseRepository(db *gorm.DB) ResponseRepository {
	return &GormResponseRepository{db: db}
}

// GormResponseRepository - реализация ResponseRepository с использованием Gorm
type GormResponseRepository struct {
	db *gorm.DB
}

func (r *GormResponseRepository) CreateResponse(response *models.Response) error {
	return r.db.Create(response).Error
}

func (r *GormResponseRepository) GetResponseByID(id uint) (*models.Response, error) {
	var response models.Response
	err := r.db.Preload("Answers").First(&response, id).Error
	if err != nil {
		return nil, err
	}
	return &response, nil
}
