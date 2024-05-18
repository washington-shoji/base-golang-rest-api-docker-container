package services

import (
	"base-golang-rest-api-docker-container/mocks"
	"base-golang-rest-api-docker-container/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTodo_Success(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := NewTodoService(mockRepo)

	todo := &models.Todo{
		Label:     "Test Todo",
		Completed: false,
	}

	mockRepo.On("CreateTodo", mock.AnythingOfType("*models.Todo")).Return(nil)

	err := todoService.CreateTodo(todo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTodo_Success(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := NewTodoService(mockRepo)

	todoID := uuid.New()
	todo := &models.Todo{
		ID:        todoID,
		Label:     "Updated Todo",
		Completed: true,
	}

	mockRepo.On("FindTodoByID", todoID).Return(todo, nil)
	mockRepo.On("UpdateTodo", mock.AnythingOfType("*models.Todo")).Return(nil)

	err := todoService.UpdateTodo(todoID, todo)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTodo_Success(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := NewTodoService(mockRepo)

	todoID := uuid.New()
	todo := &models.Todo{
		ID: todoID,
	}

	mockRepo.On("FindTodoByID", todoID).Return(todo, nil)
	mockRepo.On("DeleteTodo", todoID).Return(nil)

	err := todoService.DeleteTodo(todoID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindAllTodos_Success(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := NewTodoService(mockRepo)

	todos := []*models.Todo{
		{
			ID:        uuid.New(),
			Label:     "Todo 1",
			Completed: false,
		},
		{
			ID:        uuid.New(),
			Label:     "Todo 2",
			Completed: true,
		},
	}

	mockRepo.On("FindAllTodos").Return(todos, nil)

	result, err := todoService.FindAllTodos()

	assert.NoError(t, err)
	assert.Equal(t, todos, result)
	mockRepo.AssertExpectations(t)
}

func TestFindTodoByID_Success(t *testing.T) {
	mockRepo := new(mocks.MockTodoRepository)
	todoService := NewTodoService(mockRepo)

	todoID := uuid.New()
	todo := &models.Todo{
		ID:        todoID,
		Label:     "Test Todo",
		Completed: false,
	}

	mockRepo.On("FindTodoByID", todoID).Return(todo, nil)

	result, err := todoService.FindTodoByID(todoID)

	assert.NoError(t, err)
	assert.Equal(t, todo, result)
	mockRepo.AssertExpectations(t)
}
