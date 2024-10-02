package models

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

// Quiz представляет викторину или опрос
type Quiz struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	Questions   []Question `gorm:"foreignKey:QuizID" json:"questions,omitempty"`
}

// Question представляет вопрос в опросе
type Question struct {
	ID      uint     `gorm:"primaryKey" json:"id"`
	QuizID  uint     `gorm:"index" json:"quiz_id"`
	Text    string   `json:"text"`
	Type    string   `json:"type"`                                           // Тип вопроса: multiple_choice, single_choice, open_ended
	Options []Option `gorm:"foreignKey:QuestionID" json:"options,omitempty"` // Варианты ответа
}

// Option представляет вариант ответа на вопрос
type Option struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	QuestionID uint   `gorm:"index" json:"question_id"`
	Text       string `json:"text"`
}

// Response представляет ответы пользователя на викторину
type Response struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	QuizID    uint      `gorm:"index" json:"quiz_id"`
	CreatedAt time.Time `json:"created_at"`
	Answers   []Answer  `gorm:"foreignKey:ResponseID" json:"answers"`
}

// Answer представляет ответ на один вопрос
type Answer struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	ResponseID uint   `gorm:"index" json:"response_id"`
	QuestionID uint   `gorm:"index" json:"question_id"`
	OptionID   *uint  `gorm:"index" json:"option_id,omitempty"` // Ссылка на выбранный вариант ответа (для multiple_choice и single_choice)
	AnswerText string `json:"answer_text,omitempty"`            // Ответ на вопрос (для open_ended)
}
