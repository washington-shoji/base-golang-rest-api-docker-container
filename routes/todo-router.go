package routes

import (
	"base-golang-rest-api-docker-container/handlers"
	"base-golang-rest-api-docker-container/repositories"
	"base-golang-rest-api-docker-container/services"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitTodoRouter(e *echo.Echo, dbConn *sql.DB) {
	todoRepository := repositories.NewTodoRepositoryImpl(dbConn)
	todoService := services.NewTodoService(todoRepository)
	todoHandler := handlers.NewTodoHandler(todoService)

	group := e.Group("api/v1")

	group.GET("/todo", todoHandler.FindAllTodos)
	group.GET("/todo/:id", todoHandler.FindTodoByID)
	group.POST("/todo", todoHandler.CreateTodo)
	group.PUT("/todo/:id", todoHandler.UpdateTodo)
	group.DELETE("/todo/:id", todoHandler.DeleteTodo)
}
