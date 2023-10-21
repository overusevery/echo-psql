package psql

import (
	"database/sql"
	"overusevery/echo-psql/domain/entity"

	_ "embed"

	_ "github.com/lib/pq"
	"github.com/sony/sonyflake"
)

//go:embed sql/insert_todos.sql
var SQL_INSERT_TODOS string

//go:embed sql/select_todos_by_id.sql
var SQL_SELECT_TODOS_BY_ID string

type PSQLTodoRepository struct {
	db sql.DB
}

func NewPSQLTodoRepository(db sql.DB) *PSQLTodoRepository {
	return &PSQLTodoRepository{db: db}
}

func (r *PSQLTodoRepository) Create(todo entity.Todo) error {
	generator, err := sonyflake.New(sonyflake.Settings{})
	if err != nil {
		panic(err)
	}

	id, err := generator.NextID()
	if err != nil {
		panic(err)
	}

	// クエリを実行し、結果を取得
	_, err = r.db.Exec(SQL_INSERT_TODOS, id, todo.Content)
	if err != nil {
		panic(err)
	}

	return nil
}

func (r *PSQLTodoRepository) Get(id string) (*entity.Todo, error) {
	todo := entity.Todo{}
	err := r.db.QueryRow(SQL_SELECT_TODOS_BY_ID, id).Scan(&todo.ID, &todo.Content, &todo.Status, &todo.UpdatedAt, &todo.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}
