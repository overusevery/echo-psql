package handler

import (
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
	panic("not implemented")
}
