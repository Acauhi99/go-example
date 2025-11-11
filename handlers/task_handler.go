package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/acauhi/kanban-backend/models"
	"github.com/acauhi/kanban-backend/repository"
	"github.com/acauhi/kanban-backend/service"
)

const (
	pathTasks              = "/tasks/"
	msgInternalServerError = "Internal server error"
	msgInvalidRequestBody  = "Invalid request body"
	msgTaskIDRequired      = "Task ID is required"
	msgTaskNotFound        = "Task not found"
	msgMethodNotAllowed    = "Method not allowed"
)

type TaskHandler struct {
	service *service.TaskService
}

// NewTaskHandler cria uma nova instância do handler de tarefas
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	// TODO: implementar
	return nil
}

// ServeHTTP roteia as requisições HTTP para os handlers apropriados
func (h *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar
}

// handleCreate processa requisições POST para criar uma nova tarefa
func (h *TaskHandler) handleCreate(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar
}

// handleGetAll processa requisições GET para listar todas as tarefas
func (h *TaskHandler) handleGetAll(w http.ResponseWriter, _ *http.Request) {
	// TODO: implementar
}

// handleGetByID processa requisições GET para buscar uma tarefa por ID
func (h *TaskHandler) handleGetByID(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar
}

// handleUpdate processa requisições PUT para atualizar uma tarefa existente
func (h *TaskHandler) handleUpdate(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar
}

// handleDelete processa requisições DELETE para remover uma tarefa
func (h *TaskHandler) handleDelete(w http.ResponseWriter, r *http.Request) {
	// TODO: implementar
}
