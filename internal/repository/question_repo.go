package repository

import (
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/renanewbank/survey-api/internal/models"
)

type QuestionRepository struct {
	mu        sync.Mutex
	questions map[string]models.Question
}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{
		questions: make(map[string]models.Question),
	}
}

func (r *QuestionRepository) Create(q models.Question) models.Question {
	r.mu.Lock()
	defer r.mu.Unlock()

	q.ID = uuid.NewString()
	q.CreatedAt = time.Now()
	q.UpdatedAt = time.Now()
	q.Version = 1
	r.questions[q.ID] = q
	return q
}

func (r *QuestionRepository) GetAll() []models.Question {
	r.mu.Lock()
	defer r.mu.Unlock()

	result := make([]models.Question, 0, len(r.questions))
	for _, q := range r.questions {
		result = append(result, q)
	}
	return result
}

func (r *QuestionRepository) GetByID(id string) (models.Question, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	q, exists := r.questions[id]
	if !exists {
		return models.Question{}, errors.New("not found")
	}
	return q, nil
}

func (r *QuestionRepository) Update(id string, updated models.Question) (models.Question, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	q, exists := r.questions[id]
	if !exists {
		return models.Question{}, errors.New("not found")
	}

	updated.ID = id
	updated.CreatedAt = q.CreatedAt
	updated.UpdatedAt = time.Now()
	updated.Version = q.Version + 1
	r.questions[id] = updated

	return updated, nil
}

func (r *QuestionRepository) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.questions[id]
	if !exists {
		return errors.New("not found")
	}

	delete(r.questions, id)
	return nil
}
