package services

import (
	"base-golang-rest-api-docker-container/models"
	"base-golang-rest-api-docker-container/repositories"

	"github.com/google/uuid"
)

type TodoServiceImpl struct {
	TodoRepository repositories.TodoRepository
}

func NewTodoService(todoRepository repositories.TodoRepository) TodoService {
	return &TodoServiceImpl{
		TodoRepository: todoRepository,
	}
}

// CreateTodo implements TodoService.
func (t *TodoServiceImpl) CreateTodo(todo *models.Todo) error {
	id := uuid.New()

	todoModel := models.Todo{
		ID:        id,
		Label:     todo.Label,
		Completed: todo.Completed,
	}
	err := t.TodoRepository.CreateTodo(&todoModel)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTodo implements TodoService.
func (t *TodoServiceImpl) UpdateTodo(id uuid.UUID, todo *models.Todo) error {
	_, err := t.TodoRepository.FindTodoByID(id)
	if err != nil {
		return err
	}

	todoModel := models.Todo{
		ID:        todo.ID,
		Label:     todo.Label,
		Completed: todo.Completed,
	}

	if err := t.TodoRepository.UpdateTodo(&todoModel); err != nil {
		return err
	}

	return nil
}

// DeleteTodo implements TodoService.
func (t *TodoServiceImpl) DeleteTodo(id uuid.UUID) error {
	_, err := t.TodoRepository.FindTodoByID(id)
	if err != nil {
		return err
	}

	if err := t.TodoRepository.DeleteTodo(id); err != nil {
		return err
	}

	return nil
}

// FindAllTodos implements TodoService.
func (t *TodoServiceImpl) FindAllTodos() ([]*models.Todo, error) {
	todos, err := t.TodoRepository.FindAllTodos()
	if err != nil {
		return nil, err
	}

	return todos, nil

}

// FindTodoByID implements TodoService.
func (t *TodoServiceImpl) FindTodoByID(id uuid.UUID) (*models.Todo, error) {
	todo, err := t.TodoRepository.FindTodoByID(id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
