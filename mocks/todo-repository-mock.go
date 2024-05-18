package mocks

import (
	"base-golang-rest-api-docker-container/models"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) CreateTodo(todo *models.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoRepository) UpdateTodo(todo *models.Todo) error {
	args := m.Called(todo)
	return args.Error(0)
}

func (m *MockTodoRepository) DeleteTodo(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTodoRepository) FindAllTodos() ([]*models.Todo, error) {
	args := m.Called()
	return args.Get(0).([]*models.Todo), args.Error(1)
}

func (m *MockTodoRepository) FindTodoByID(id uuid.UUID) (*models.Todo, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Todo), args.Error(1)
}
