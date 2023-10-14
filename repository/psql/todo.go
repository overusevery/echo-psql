package psql

import (
	"database/sql"
	"overusevery/echo-psql/domain/entity"

	_ "github.com/lib/pq"
	"github.com/sony/sonyflake"
)

type PSQLTodoRepository struct {
	db sql.DB
}

func NewPSQLTodoRepository() *PSQLTodoRepository {
	return &PSQLTodoRepository{}
}

func (r *PSQLTodoRepository) Create(todo entity.Todo) error {
	// PostgreSQLへの接続情報
	connStr := "user=root dbname=mydb password=changeme sslmode=disable"

	// PostgreSQLデータベースに接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	generator, err := sonyflake.New(sonyflake.Settings{})
	if err != nil {
		panic(err)
	}

	id, err := generator.NextID()
	if err != nil {
		panic(err)
	}

	// クエリを実行し、結果を取得
	_, err = db.Exec("INSERT INTO todos (ID, Content, Status, UpdatedAt, CreatedAt) VALUES ($1, $2, true, NOW(), NOW());", id, todo.Content)
	if err != nil {
		panic(err)
	}

	return nil
}
