package test

import (
	"context"
	"reflect"

	"InnoTaxi/internal/pkg/dto"
	"InnoTaxi/internal/pkg/model"

	"github.com/golang/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockIUserRepository) AddUser(ctx context.Context, user model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockIUserRepositoryMockRecorder) AddUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockIUserRepository)(nil).AddUser), ctx, user)
}

// ChangeUser mocks base method.
func (m *MockIUserRepository) UpdateUser(ctx context.Context, request dto.ChangeUserRequest, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, request, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUser indicates an expected call of ChangeUser.
func (mr *MockIUserRepositoryMockRecorder) ChangeUser(ctx, request, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserRepository)(nil).UpdateUser), ctx, request, id)
}

// DeleteUser mocks base method.
func (m *MockIUserRepository) DeleteUser(context context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", context, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserRepositoryMockRecorder) DeleteUser(context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserRepository)(nil).DeleteUser), context, id)
}

// DoesPhoneExist mocks base method.
func (m *MockIUserRepository) DoesPhoneExist(ctx context.Context, phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesPhoneExist", ctx, phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// DoesPhoneExist indicates an expected call of DoesPhoneExist.
func (mr *MockIUserRepositoryMockRecorder) DoesPhoneExist(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesPhoneExist", reflect.TypeOf((*MockIUserRepository)(nil).DoesPhoneExist), ctx, phone)
}

// GetAllUsers mocks base method.
func (m *MockIUserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserRepositoryMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserRepository)(nil).GetAllUsers), ctx)
}

// GetUserById mocks base method.
func (m *MockIUserRepository) GetUserById(ctx context.Context, id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockIUserRepositoryMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockIUserRepository)(nil).GetUserById), ctx, id)
}

// GetUserByPhone mocks base method.
func (m *MockIUserRepository) GetUserByPhone(ctx context.Context, phone string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhone", ctx, phone)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhone indicates an expected call of GetUserByPhone.
func (mr *MockIUserRepositoryMockRecorder) GetUserByPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhone", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByPhone), ctx, phone)
}
