package services

import (
	"base-golang-rest-api-docker-container/models"

	"github.com/google/uuid"
)

type TodoService interface {
	CreateTodo(todo *models.Todo) error
	UpdateTodo(id uuid.UUID, model *models.Todo) error
	DeleteTodo(id uuid.UUID) error
	FindTodoByID(id uuid.UUID) (*models.Todo, error)
	FindAllTodos() ([]*models.Todo, error)
}
