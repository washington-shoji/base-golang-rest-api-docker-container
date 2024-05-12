package repositories

import (
	"base-golang-rest-api-docker-container/models"

	"github.com/google/uuid"
)

type TodoRepository interface {
	CreateTodo(todo *models.Todo) error
	UpdateTodo(todo *models.Todo) error
	DeleteTodo(id uuid.UUID) error
	FindTodoByID(id uuid.UUID) (*models.Todo, error)
	FindAllTodos() ([]*models.Todo, error)
}
