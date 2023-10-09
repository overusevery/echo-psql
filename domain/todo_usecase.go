package domain

import (
	"overusevery/echo-psql/domain/entity"
)

type TodoUsecase struct {
	todoRepository TodoRepository
}

func (tu *TodoUsecase) Create(content string) error {
	newTodo := entity.Todo{
		Content: content,
		Status:  false,
	}
	return tu.todoRepository.Create(newTodo)
}

type TodoRepository interface {
	Create(todo entity.Todo) error
}
