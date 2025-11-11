package service

import (
	"errors"

	"github.com/acauhi/kanban-backend/models"
	"github.com/acauhi/kanban-backend/repository"
	"github.com/google/uuid"
)

var (
	ErrInvalidTitle  = errors.New("title is required")
	ErrInvalidStatus = errors.New("invalid status")
)

type TaskService struct {
	repo repository.TaskRepository
}

// NewTaskService cria uma nova instância do serviço de tarefas
func NewTaskService(repo repository.TaskRepository) *TaskService {
	// TODO: implementar
	return nil
}

// CreateTask cria uma nova tarefa com status inicial "todo"
func (s *TaskService) CreateTask(req models.CreateTaskRequest) (*models.Task, error) {
	// TODO: implementar
	return nil, nil
}

// GetAllTasks retorna todas as tarefas cadastradas
func (s *TaskService) GetAllTasks() ([]*models.Task, error) {
	// TODO: implementar
	return nil, nil
}

// GetTaskByID busca uma tarefa específica pelo ID
func (s *TaskService) GetTaskByID(id string) (*models.Task, error) {
	// TODO: implementar
	return nil, nil
}

// UpdateTask atualiza os campos de uma tarefa existente
func (s *TaskService) UpdateTask(id string, req models.UpdateTaskRequest) (*models.Task, error) {
	// TODO: implementar
	return nil, nil
}

// DeleteTask remove uma tarefa pelo ID
func (s *TaskService) DeleteTask(id string) error {
	// TODO: implementar
	return nil
}

// isValidStatus valida se o status fornecido é um dos valores permitidos
func isValidStatus(status models.Status) bool {
	// TODO: implementar
	return false
}
