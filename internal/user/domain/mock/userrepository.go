// Code generated by MockGen. DO NOT EDIT.
// Source: internal/user/domain/interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	postgres "github.com/cyworld8x/go-postgres-kubernetes-grpc/internal/user/infrastructure/repository/postgres"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// ChangeUserPassword mocks base method.
func (m *MockUserRepository) ChangeUserPassword(ctx context.Context, arg postgres.ChangeUserPasswordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUserPassword", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword.
func (mr *MockUserRepositoryMockRecorder) ChangeUserPassword(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockUserRepository)(nil).ChangeUserPassword), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(ctx context.Context, arg postgres.CreateUserParams) (postgres.DbUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(postgres.DbUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), ctx, arg)
}

// DeactiveUser mocks base method.
func (m *MockUserRepository) DeactiveUser(ctx context.Context, id uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeactiveUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeactiveUser indicates an expected call of DeactiveUser.
func (mr *MockUserRepositoryMockRecorder) DeactiveUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeactiveUser", reflect.TypeOf((*MockUserRepository)(nil).DeactiveUser), ctx, id)
}

// GetLogin mocks base method.
func (m *MockUserRepository) GetLogin(ctx context.Context, username string) (postgres.DbUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogin", ctx, username)
	ret0, _ := ret[0].(postgres.DbUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogin indicates an expected call of GetLogin.
func (mr *MockUserRepositoryMockRecorder) GetLogin(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogin", reflect.TypeOf((*MockUserRepository)(nil).GetLogin), ctx, username)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(ctx context.Context, id uuid.UUID) (postgres.DbUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(postgres.DbUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), ctx, id)
}

// ListUsers mocks base method.
func (m *MockUserRepository) ListUsers(ctx context.Context) ([]postgres.DbUser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx)
	ret0, _ := ret[0].([]postgres.DbUser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockUserRepositoryMockRecorder) ListUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockUserRepository)(nil).ListUsers), ctx)
}

// UpdateUserInfo mocks base method.
func (m *MockUserRepository) UpdateUserInfo(ctx context.Context, arg postgres.UpdateUserInfoParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserInfo", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserInfo indicates an expected call of UpdateUserInfo.
func (mr *MockUserRepositoryMockRecorder) UpdateUserInfo(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserInfo", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserInfo), ctx, arg)
}