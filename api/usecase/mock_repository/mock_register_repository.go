// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/repository/register_repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "api/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRegisterUserRepository is a mock of RegisterUserRepository interface.
type MockRegisterUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterUserRepositoryMockRecorder
}

// MockRegisterUserRepositoryMockRecorder is the mock recorder for MockRegisterUserRepository.
type MockRegisterUserRepositoryMockRecorder struct {
	mock *MockRegisterUserRepository
}

// NewMockRegisterUserRepository creates a new mock instance.
func NewMockRegisterUserRepository(ctrl *gomock.Controller) *MockRegisterUserRepository {
	mock := &MockRegisterUserRepository{ctrl: ctrl}
	mock.recorder = &MockRegisterUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegisterUserRepository) EXPECT() *MockRegisterUserRepositoryMockRecorder {
	return m.recorder
}

// InsertData mocks base method.
func (m *MockRegisterUserRepository) InsertData(user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertData", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertData indicates an expected call of InsertData.
func (mr *MockRegisterUserRepositoryMockRecorder) InsertData(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertData", reflect.TypeOf((*MockRegisterUserRepository)(nil).InsertData), user)
}
