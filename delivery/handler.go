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
	e.POST("/todos", handler.Create)
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
