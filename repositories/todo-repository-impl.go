package repositories

import (
	"base-golang-rest-api-docker-container/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type TodoRepositoryImpl struct {
	DB *sql.DB
}

func NewTodoRepositoryImpl(db *sql.DB) TodoRepository {
	return &TodoRepositoryImpl{
		DB: db,
	}
}

// CreateTodo implements TodoRepository.
func (t *TodoRepositoryImpl) CreateTodo(todo *models.Todo) error {
	query := `
	INSERT INTO todo (id, label, completed)
	VALUES ($1, $2, $3)
	`
	_, err := t.DB.Query(
		query,
		todo.ID,
		todo.Label,
		todo.Completed,
	)
	if err != nil {
		return err
	}
	return nil
}

// UpdateTodo implements TodoRepository.
func (t *TodoRepositoryImpl) UpdateTodo(todo *models.Todo) error {
	query := `
	UPDATE todo
	SET label = $2, completed = $3
	WHERE id = $1
	`
	_, err := t.DB.Query(
		query,
		todo.ID,
		todo.Label,
		todo.Completed,
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteTodo implements TodoRepository.
func (t *TodoRepositoryImpl) DeleteTodo(id uuid.UUID) error {
	query := `
	DELETE FROM todo
	WHERE id = $1
	`
	_, err := t.DB.Query(
		query,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

// FindAllTodos implements TodoRepository.
func (t *TodoRepositoryImpl) FindAllTodos() ([]*models.Todo, error) {
	query := `
	SELECT * FROM todo
	`
	rows, err := t.DB.Query(query)
	if err != nil {
		return nil, err
	}

	todos := []*models.Todo{}
	for rows.Next() {
		todo, err := scanIntoTodo(rows)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// FindTodoByID implements TodoRepository.
func (t *TodoRepositoryImpl) FindTodoByID(id uuid.UUID) (*models.Todo, error) {
	query := `
	SELECT * FROM todo
	WHERE id = $1
	`

	rows, err := t.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoTodo(rows)
	}

	return nil, fmt.Errorf("todo %s not found", id)
}

func scanIntoTodo(rows *sql.Rows) (*models.Todo, error) {
	todo := &models.Todo{}
	err := rows.Scan(
		&todo.ID,
		&todo.Label,
		&todo.Completed,
	)

	return todo, err
}
