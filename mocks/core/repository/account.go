// Code generated by MockGen. DO NOT EDIT.
// Source: account.go
//
// Generated by this command:
//
//	mockgen -source=account.go -destination=../../../mocks/core/repository/account.go -package=repository
//

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"
	request "routines/api/handlers/contract/request"

	gomock "go.uber.org/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
	isgomock struct{}
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAccountRepository) Create(ctx context.Context, request *request.AccountRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAccountRepositoryMockRecorder) Create(ctx, request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAccountRepository)(nil).Create), ctx, request)
}