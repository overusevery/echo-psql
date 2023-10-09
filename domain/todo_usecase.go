package domain

import (
	"errors"
	"overusevery/echo-psql/domain/entity"
)

type TodoUsecase struct {
	todoRepository TodoRepository
}

const MAX_CONTENT_LENGTH = 400

func (tu *TodoUsecase) Create(content string) error {
	if len(content) > MAX_CONTENT_LENGTH {
		return errors.New("content is too long")
	}
	newTodo := entity.Todo{
		Content: content,
		Status:  false,
	}
	return tu.todoRepository.Create(newTodo)
}

type TodoRepository interface {
	Create(todo entity.Todo) error
}
