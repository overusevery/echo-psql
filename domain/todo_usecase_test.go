package domain

import (
	"errors"
	mock_domain "overusevery/echo-psql/repository/mock"
	"testing"

	"go.uber.org/mock/gomock"
)

const contents_with_400_string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
const contents_with_401_string = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

func TestTodoUsecase_Create_Success(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Success to Create Todo", args: args{content: "write draft"}, wantErr: false},
		{name: "Fail to Create When content's length > 400:len=400", args: args{content: contents_with_400_string}, wantErr: false},
		{name: "Fail to Create When content's length > 400:len=401", args: args{content: contents_with_401_string}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mock := mock_domain.NewMockTodoRepository(ctrl)
			mock.
				EXPECT().
				Create(gomock.Any()).
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
