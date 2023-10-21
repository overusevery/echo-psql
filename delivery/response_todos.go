package handler

import "overusevery/echo-psql/domain/entity"

type ResponseTodos struct {
	entity.Todo
}

func Todo2ResponseTodos(todo entity.Todo) ResponseTodos {
	return ResponseTodos{
		Todo: todo,
	}
}
