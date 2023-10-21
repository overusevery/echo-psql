package handler

import (
	"net/http"
	"overusevery/echo-psql/domain"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	usecase domain.TodoUsecase
}

func TodoRouter(e *echo.Echo, r domain.TodoRepository) {
	usecase := domain.NewTodoUsecase(r)
	handler := &TodoHandler{
		usecase: *usecase,
	}
	e.GET("/todos/:id", handler.Get)
	e.POST("/todos", handler.Create)
}

func (h *TodoHandler) Get(c echo.Context) error {
	id := c.Param("id")
	todo, err := h.usecase.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, Todo2ResponseTodos(*todo))
}

func (h *TodoHandler) Create(c echo.Context) error {
	req := new(RequestTodosCreate)
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := h.usecase.Create(req.Content); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, req)
}
