// Code generated by MockGen. DO NOT EDIT.
// Source: base_db.go
//
// Generated by this command:
//
//	mockgen -source=base_db.go -destination=../mocks/db/base_db.go -package=db
//

// Package db is a generated GoMock package.
package db

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBaseDB is a mock of BaseDB interface.
type MockBaseDB struct {
	ctrl     *gomock.Controller
	recorder *MockBaseDBMockRecorder
	isgomock struct{}
}

// MockBaseDBMockRecorder is the mock recorder for MockBaseDB.
type MockBaseDBMockRecorder struct {
	mock *MockBaseDB
}

// NewMockBaseDB creates a new mock instance.
func NewMockBaseDB(ctrl *gomock.Controller) *MockBaseDB {
	mock := &MockBaseDB{ctrl: ctrl}
	mock.recorder = &MockBaseDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBaseDB) EXPECT() *MockBaseDBMockRecorder {
	return m.recorder
}

// DB mocks base method.
func (m *MockBaseDB) DB(ctx context.Context) *sql.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB", ctx)
	ret0, _ := ret[0].(*sql.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockBaseDBMockRecorder) DB(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockBaseDB)(nil).DB), ctx)
}

// Insert mocks base method.
func (m *MockBaseDB) Insert(ctx context.Context, query string, data ...any) error {
	m.ctrl.T.Helper()
	varargs := []any{ctx, query}
	for _, a := range data {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockBaseDBMockRecorder) Insert(ctx, query any, data ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, query}, data...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockBaseDB)(nil).Insert), varargs...)
}
