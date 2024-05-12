package handlers

import (
	"base-golang-rest-api-docker-container/models"
	"base-golang-rest-api-docker-container/services"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	TodoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{
		TodoService: todoService,
	}
}

func (h *TodoHandler) CreateTodo(c echo.Context) error {
	todoModel := models.Todo{}
	err := c.Bind(&todoModel)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if err := h.TodoService.CreateTodo(&todoModel); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Todo Created!")
}

func (h *TodoHandler) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	todoModel := models.Todo{}
	err := c.Bind(&todoModel)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := h.TodoService.UpdateTodo(uid, &todoModel); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Todo Updated!")
}

func (h *TodoHandler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	if err := h.TodoService.DeleteTodo(uid); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, "Todo Deleted!")
}

func (h *TodoHandler) FindTodoByID(c echo.Context) error {
	id := c.Param("id")

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	todo, err := h.TodoService.FindTodoByID(uid)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, todo)
}

func (h *TodoHandler) FindAllTodos(c echo.Context) error {

	todos, err := h.TodoService.FindAllTodos()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, todos)
}
