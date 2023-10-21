package psql

import (
	"database/sql"
	"overusevery/echo-psql/domain/entity"
	"testing"
	"time"

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
			//setup
			db := setuptestDB()
			defer db.Close()
			r := &PSQLTodoRepository{db: *db}
			EXECUTE_PSQL("DELETE FROM public.todos WHERE content = 'ToDo xxxxxyyyyzzz' and status = true;")

			//test
			if err := r.Create(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("PSQLTodoRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !CHECK_WITH_SELECT_PSQL("SELECT count(1) = 1 FROM public.todos WHERE content = 'ToDo xxxxxyyyyzzz' and status = true;") {
				t.Errorf("Record not found.Failed to Create records")
			}
		})
	}
}

func TestPSQLTodoRepository_Create_twice(t *testing.T) {
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
			name: "create same contents twice will generate 2 records",
			args: args{
				todo: entity.Todo{
					Content: "ToDo twice",
					Status:  false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//setup
			db := setuptestDB()
			defer db.Close()
			r := &PSQLTodoRepository{db: *db}
			EXECUTE_PSQL("DELETE FROM public.todos WHERE content = 'ToDo twice' and status = true;")

			//test
			if err := r.Create(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("PSQLTodoRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := r.Create(tt.args.todo); (err != nil) != tt.wantErr {
				t.Errorf("PSQLTodoRepository.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !CHECK_WITH_SELECT_PSQL("SELECT count(1) = 2 FROM public.todos WHERE content = 'ToDo twice' and status = true;") {
				t.Errorf("Record not found. Failed to Create records")
			}
		})
	}
}
func TestPSQLTodoRepository_Get(t *testing.T) {
	//setup
	db := setuptestDB()
	defer db.Close()
	r := &PSQLTodoRepository{db: *db}
	EXECUTE_PSQL("DELETE FROM public.todos WHERE id = '1234' and status = true;")
	EXECUTE_PSQL("INSERT INTO public.todos (id, content, status, updatedat, createdat) VALUES ('1234', 'contents for xxx', true, '2004-10-20 10:23:54', '2004-10-19 10:23:54');")

	//test
	actualTodoGet, err := r.Get("1234")
	if err != nil {
		t.Errorf("PSQLTodoRepository.Get() error = %v", err)
	}
	if actualTodoGet.ID != 1234 {
		t.Errorf("PSQLTodoRepository.Get(); ID doesn't match ; actual = %v, expected = %v", actualTodoGet.ID, 1234)
	}
	if actualTodoGet.Content != "contents for xxx" {
		t.Errorf("PSQLTodoRepository.Get(); Content doesn't match ; actual = %v, expected = %v", actualTodoGet.Content, "contents for xxx")
	}
	if actualTodoGet.Status != true {
		t.Errorf("PSQLTodoRepository.Get(); Status doesn't match ; actual = %v, expected = %v", actualTodoGet.Status, true)
	}
	expectedUpdatedate, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2004-10-20 10:23:54")
	if actualTodoGet.UpdatedAt == expectedUpdatedate {
		t.Errorf("PSQLTodoRepository.Get(); UpdatedAt doesn't match ; actual = %v, expected = %v", actualTodoGet.UpdatedAt, expectedUpdatedate)
	}
	expectedCreateddate, _ := time.Parse("2006-01-02T15:04:05Z07:00", "2004-10-19 10:23:54")
	if actualTodoGet.CreatedAt == expectedCreateddate {
		t.Errorf("PSQLTodoRepository.Get(); CreatedAt doesn't match ; actual = %v, expected = %v", actualTodoGet.CreatedAt, expectedCreateddate)
	}
}

func setuptestDB() *sql.DB {
	// PostgreSQLへの接続情報
	connStr := "user=root dbname=mydb password=changeme sslmode=disable"

	// PostgreSQLデータベースに接続
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func EXECUTE_PSQL(query string) {
	db := setuptestDB()
	defer db.Close()

	// クエリを実行し、結果を取得
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func CHECK_WITH_SELECT_PSQL(query string) bool {
	db := setuptestDB()
	defer db.Close()

	// クエリを実行し、結果を取得
	row := db.QueryRow(query)
	var isTrue bool
	err := row.Scan(&isTrue)
	if err != nil {
		panic(err)
	}
	return isTrue
}
