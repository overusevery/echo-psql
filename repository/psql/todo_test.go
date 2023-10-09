package psql

import (
	"database/sql"
	"fmt"
	"overusevery/echo-psql/domain/entity"
	"testing"

	_ "github.com/lib/pq"
)

func TestPSQLTodoRepository_Create(t *testing.T) {
	type args struct {
		todo entity.Todo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Success to create",
			args: args{
				todo: entity.Todo{
					Content: "ToDo xxxxxyyyyzzz",
					Status:  false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &PSQLTodoRepository{}
			if err := r.Create(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("PSQLTodoRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !SELECT_PSQL("SELECT count(1) = 1 FROM public.todos WHERE content = 'ToDo xxxxxyyyyzzz' and status = true;") {
				t.Errorf("Record not found.Failed to Create records")
			}
		})
	}
}

func SELECT_PSQL(query string) bool {
	// PostgreSQLへの接続情報
	connStr := "user=root dbname=mydb password=changeme sslmode=disable"

	// PostgreSQLデータベースに接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// クエリを実行し、結果を取得
	row := db.QueryRow(query)
	var isTrue bool
	err = row.Scan(&isTrue)
	if err != nil {
		panic(err)
	}
	fmt.Printf("isTure: %v", isTrue)
	return isTrue
}
