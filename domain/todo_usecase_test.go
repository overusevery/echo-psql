package domain

import (
	"errors"
	mock_domain "overusevery/echo-psql/repository/mock"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestTodoUsecase_Create_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := mock_domain.NewMockTodoRepository(ctrl)
	mock.
		EXPECT().
		Create(gomock.Any()).
		Return(nil)

	tu := &TodoUsecase{
		todoRepository: mock,
	}
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success to Create Todo", args: args{content: "write draft"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tu.Create(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoUsecase_Create_RepoError(t *testing.T) {
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
			if err := tu.Create(tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("TodoUsecase.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
