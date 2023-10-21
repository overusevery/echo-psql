package domain

import (
	"errors"
	"overusevery/echo-psql/domain/entity"
	mock_domain "overusevery/echo-psql/repository/mock"
	"reflect"
	"testing"
	"time"

	"go.uber.org/mock/gomock"
)

const contents_with_400_string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
const contents_with_401_string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

func TestTodoUsecase_Create_Success(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name            string
		args            args
		timesToCallRepo int
		wantErr         bool
	}{
		{name: "Success to Create Todo", args: args{content: "write draft"}, timesToCallRepo: 1, wantErr: false},
		{name: "Fail to Create When content's length > 400:len=400", args: args{content: contents_with_400_string}, timesToCallRepo: 1, wantErr: false},
		{name: "Fail to Create When content's length > 400:len=401", args: args{content: contents_with_401_string}, timesToCallRepo: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mock := mock_domain.NewMockTodoRepository(ctrl)
			mock.
				EXPECT().
				Create(gomock.Any()).
				Times(tt.timesToCallRepo).
				Return(nil)

			tu := &TodoUsecase{
				todoRepository: mock,
			}
			if err := tu.Create(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoUsecase_Create_RepoError(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Failed to Create Todo When Repo returns error", args: args{content: "write draft"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mock := mock_domain.NewMockTodoRepository(ctrl)
			errorFromMock := errors.New("something went wrong when calling repository")
			mock.
				EXPECT().
				Create(gomock.Any()).
				Return(errorFromMock)

			tu := &TodoUsecase{
				todoRepository: mock,
			}
			if err := tu.Create(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoUsecase_Get(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success to Get Todo", args: args{id: "111"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mock := mock_domain.NewMockTodoRepository(ctrl)
			sampleTodo := entity.Todo{
				ID:        111,
				Content:   "content111",
				Status:    false,
				UpdatedAt: time.Date(2022, 5, 1, 9, 0, 0, 0, time.Local),
				CreatedAt: time.Date(2022, 4, 1, 9, 0, 0, 0, time.Local),
			}
			mock.
				EXPECT().
				Get(gomock.Any()).
				Return(sampleTodo, nil)

			tu := &TodoUsecase{
				todoRepository: mock,
			}
			todoActual, err := tu.Get(tt.args.id)
			if err != nil && !tt.wantErr {
				t.Errorf("TodoUsecase.Get(), test name = %v , error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if !reflect.DeepEqual(todoActual, sampleTodo) {
				t.Errorf("TodoUsecase.Get(), test name = %v , data should not be modified, actual = %v, expected = %v", tt.name, todoActual, sampleTodo)
			}
		})
	}
}
