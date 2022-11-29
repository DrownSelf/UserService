package test

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"

	"github.com/DrownSelf/UserService/internal/entities"
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
func (m *MockIUserRepository) AddUser(ctx context.Context, user entities.User) (int, error) {
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

// AppendRating mocks base method.
func (m *MockIUserRepository) AppendRating(ctx context.Context, id int, rating float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppendRating", ctx, id, rating)
	ret0, _ := ret[0].(error)
	return ret0
}

// AppendRating indicates an expected call of AppendRating.
func (mr *MockIUserRepositoryMockRecorder) AppendRating(ctx, id, rating interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppendRating", reflect.TypeOf((*MockIUserRepository)(nil).AppendRating), ctx, id, rating)
}

// DeleteUser mocks base method.
func (m *MockIUserRepository) DeleteUser(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserRepositoryMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserRepository)(nil).DeleteUser), ctx, id)
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
func (m *MockIUserRepository) GetAllUsers(ctx context.Context) ([]entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", ctx)
	ret0, _ := ret[0].([]entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockIUserRepositoryMockRecorder) GetAllUsers(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockIUserRepository)(nil).GetAllUsers), ctx)
}

// GetAverageRating mocks base method.
func (m *MockIUserRepository) GetAverageRating(ctx context.Context, id int) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAverageRating", ctx, id)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAverageRating indicates an expected call of GetAverageRating.
func (mr *MockIUserRepositoryMockRecorder) GetAverageRating(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAverageRating", reflect.TypeOf((*MockIUserRepository)(nil).GetAverageRating), ctx, id)
}

// GetUserById mocks base method.
func (m *MockIUserRepository) GetUserById(ctx context.Context, id int) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserById", ctx, id)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById.
func (mr *MockIUserRepositoryMockRecorder) GetUserById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockIUserRepository)(nil).GetUserById), ctx, id)
}

// GetUserByPhone mocks base method.
func (m *MockIUserRepository) GetUserByPhone(ctx context.Context, phone string) (entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByPhone", ctx, phone)
	ret0, _ := ret[0].(entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByPhone indicates an expected call of GetUserByPhone.
func (mr *MockIUserRepositoryMockRecorder) GetUserByPhone(ctx, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByPhone", reflect.TypeOf((*MockIUserRepository)(nil).GetUserByPhone), ctx, phone)
}

// RelateRating mocks base method.
func (m *MockIUserRepository) RelateRating(ctx context.Context, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelateRating", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelateRating indicates an expected call of RelateRating.
func (mr *MockIUserRepositoryMockRecorder) RelateRating(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelateRating", reflect.TypeOf((*MockIUserRepository)(nil).RelateRating), ctx, id)
}

// UpdateUser mocks base method.
func (m *MockIUserRepository) UpdateUser(ctx context.Context, user entities.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserRepositoryMockRecorder) UpdateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserRepository)(nil).UpdateUser), ctx, user)
}
