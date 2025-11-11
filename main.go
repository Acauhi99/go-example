package main

import (
	"log"
	"net/http"

	"github.com/acauhi/kanban-backend/handlers"
	"github.com/acauhi/kanban-backend/repository"
	"github.com/acauhi/kanban-backend/service"
)

// main inicializa o servidor HTTP com todas as dependências
func main() {
	// TODO: implementar
}

// corsMiddleware adiciona headers CORS para permitir requisições do frontend
func corsMiddleware(next http.Handler) http.Handler {
	// TODO: implementar
	return nil
}
