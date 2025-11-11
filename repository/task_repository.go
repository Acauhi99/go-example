package repository

import (
	"errors"
	"sync"

	"github.com/acauhi/kanban-backend/models"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	Create(task *models.Task) error
	GetAll() ([]*models.Task, error)
	GetByID(id string) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id string) error
}

type InMemoryTaskRepository struct {
	tasks map[string]*models.Task
	mu    sync.RWMutex
}

// NewInMemoryTaskRepository cria uma nova instância do repositório em memória
func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	// TODO: implementar
	return nil
}

// Create adiciona uma nova tarefa ao repositório
func (r *InMemoryTaskRepository) Create(task *models.Task) error {
	// TODO: implementar
	return nil
}

// GetAll retorna todas as tarefas armazenadas
func (r *InMemoryTaskRepository) GetAll() ([]*models.Task, error) {
	// TODO: implementar
	return nil, nil
}

// GetByID busca uma tarefa específica pelo ID
func (r *InMemoryTaskRepository) GetByID(id string) (*models.Task, error) {
	// TODO: implementar
	return nil, nil
}

// Update atualiza uma tarefa existente no repositório
func (r *InMemoryTaskRepository) Update(task *models.Task) error {
	// TODO: implementar
	return nil
}

// Delete remove uma tarefa do repositório pelo ID
func (r *InMemoryTaskRepository) Delete(id string) error {
	// TODO: implementar
	return nil
}
