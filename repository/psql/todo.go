package psql

import (
	"database/sql"
	"fmt"
	"overusevery/echo-psql/domain/entity"

	_ "github.com/lib/pq"
)

type PSQLTodoRepository struct {
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

	// クエリを実行し、結果を取得
	rows, err := db.Query("SELECT id FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	fmt.Printf("hello")

	// 結果を処理
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		fmt.Printf("ID: %v", id)
	}

	// エラーチェック
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return nil
}
