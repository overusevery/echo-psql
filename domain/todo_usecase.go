package domain

import (
	"errors"
	"overusevery/echo-psql/domain/entity"
)

type TodoUsecase struct {
	todoRepository TodoRepository
}

func NewTodoUsecase(r TodoRepository) *TodoUsecase {
	return &TodoUsecase{todoRepository: r}
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

func (tu *TodoUsecase) Get(id string) (entity.Todo, error) {
	todo, err := tu.todoRepository.Get(id)
	return todo, err
}

type TodoRepository interface {
	Create(todo entity.Todo) error
	Get(id string) (entity.Todo, error)
}
