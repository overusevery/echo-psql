// Code generated by MockGen. DO NOT EDIT.
// Source: domain/todo_usecase.go
//
// Generated by this command:
//
//	mockgen -source domain/todo_usecase.go -destination repository/mock/todo.go
//
// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	entity "overusevery/echo-psql/domain/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTodoRepository is a mock of TodoRepository interface.
type MockTodoRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTodoRepositoryMockRecorder
}

// MockTodoRepositoryMockRecorder is the mock recorder for MockTodoRepository.
type MockTodoRepositoryMockRecorder struct {
	mock *MockTodoRepository
}

// NewMockTodoRepository creates a new mock instance.
func NewMockTodoRepository(ctrl *gomock.Controller) *MockTodoRepository {
	mock := &MockTodoRepository{ctrl: ctrl}
	mock.recorder = &MockTodoRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTodoRepository) EXPECT() *MockTodoRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTodoRepository) Create(todo entity.Todo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", todo)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockTodoRepositoryMockRecorder) Create(todo any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTodoRepository)(nil).Create), todo)
}
