package integration

import (
	"base-golang-rest-api-docker-container/models"
	"base-golang-rest-api-docker-container/repositories"
	"base-golang-rest-api-docker-container/services"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func setup() {
	var err error
	connStr := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", "postgres", "postgres", "postgres",
	)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to the test database: %v", err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS todo (
		id uuid PRIMARY KEY NOT NULL,
		label VARCHAR(50) UNIQUE NOT NULL,
		completed BOOLEAN NOT NULL DEFAULT FALSE
	)
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Fatalf("Failed to execute schema: %v", err)
	}

}

func teardown() {
	if db != nil {
		db.Close()
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestTodoIntegration(t *testing.T) {
	todoRepo := repositories.NewTodoRepositoryImpl(db)
	todoService := services.NewTodoService(todoRepo)

	t.Run("CreateTodo", func(t *testing.T) {
		todo := &models.Todo{
			Label:     "Integration Test Todo",
			Completed: false,
		}

		err := todoService.CreateTodo(todo)
		assert.NoError(t, err)

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM todo WHERE label=$1", "Integration Test Todo").Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 1, count)
	})

	t.Run("UpdateTodo", func(t *testing.T) {
		todoID := uuid.New()
		_, err := db.Exec("INSERT INTO todo (id, label, completed) VALUES ($1, $2, $3)", todoID, "Old Label", false)
		assert.NoError(t, err)

		updatedTodo := &models.Todo{
			ID:        todoID,
			Label:     "Updated Label",
			Completed: true,
		}

		err = todoService.UpdateTodo(todoID, updatedTodo)
		assert.NoError(t, err)

		var label string
		var completed bool
		err = db.QueryRow("SELECT label, completed FROM todo WHERE id=$1", todoID).Scan(&label, &completed)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Label", label)
		assert.Equal(t, true, completed)
	})

	t.Run("DeleteTodo", func(t *testing.T) {
		todoID := uuid.New()
		_, err := db.Exec("INSERT INTO todo (id, label, completed) VALUES ($1, $2, $3)", todoID, "Delete Me", false)
		assert.NoError(t, err)

		err = todoService.DeleteTodo(todoID)
		assert.NoError(t, err)

		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM todo WHERE id=$1", todoID).Scan(&count)
		assert.NoError(t, err)
		assert.Equal(t, 0, count)
	})

	t.Run("FindAllTodos", func(t *testing.T) {
		db.Exec("DELETE * FROM todo")

		_, err := db.Exec("INSERT INTO todo (id, label, completed) VALUES ($1, $2, $3)", uuid.New(), "Todo 1", false)
		assert.NoError(t, err)
		_, err = db.Exec("INSERT INTO todo (id, label, completed) VALUES ($1, $2, $3)", uuid.New(), "Todo 2", true)
		assert.NoError(t, err)

		todos, err := todoService.FindAllTodos()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(todos))
	})

	t.Run("FindTodoByID", func(t *testing.T) {
		todoID := uuid.New()
		_, err := db.Exec("INSERT INTO todo (id, label, completed) VALUES ($1, $2, $3)", todoID, "Find Me", false)
		assert.NoError(t, err)

		todo, err := todoService.FindTodoByID(todoID)
		assert.NoError(t, err)
		assert.Equal(t, "Find Me", todo.Label)
		assert.Equal(t, false, todo.Completed)
	})
}
