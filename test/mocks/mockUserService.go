package test

import (
	"context"
	"reflect"

	"InnoTaxi/internal/pkg/dto"
	"InnoTaxi/internal/pkg/model"

	"github.com/golang/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// ChangeUserPassword mocks base method.
func (m *MockIUserService) UpdateUser(ctx context.Context, request dto.ChangeUserRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword.
func (mr *MockIUserServiceMockRecorder) ChangeUserPassword(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserService)(nil).UpdateUser), ctx, request)
}

// DeleteUser mocks base method.
func (m *MockIUserService) DeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserServiceMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserService)(nil).DeleteUser), ctx, id)
}

// GetUserByPhone mocks base method.
func (m *MockIUserService) GetUserByPhone(ctx context.Context, phoneNumber string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhone", ctx, phoneNumber)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhone indicates an expected call of GetUserByPhone.
func (mr *MockIUserServiceMockRecorder) GetUserByPhone(ctx, phoneNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhone", reflect.TypeOf((*MockIUserService)(nil).GetUserByPhone), ctx, phoneNumber)
}

// LogInUser mocks base method.
func (m *MockIUserService) LogInUser(ctx context.Context, request dto.LogInUserRequest) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogInUser", ctx, request)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogInUser indicates an expected call of LogInUser.
func (mr *MockIUserServiceMockRecorder) LogInUser(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogInUser", reflect.TypeOf((*MockIUserService)(nil).LogInUser), ctx, request)
}

// LogOutUser mocks base method.
func (m *MockIUserService) LogOutUser(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogOutUser", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogOutUser indicates an expected call of LogOutUser.
func (mr *MockIUserServiceMockRecorder) LogOutUser(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogOutUser", reflect.TypeOf((*MockIUserService)(nil).LogOutUser), ctx, token)
}

// RegisterUser mocks base method.
func (m *MockIUserService) RegisterUser(ctx context.Context, user dto.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockIUserServiceMockRecorder) RegisterUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserService)(nil).RegisterUser), ctx, user)
}
