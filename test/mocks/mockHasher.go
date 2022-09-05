package test

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockIHasher is a mock of IHasher interface.
type MockIHasher struct {
	ctrl     *gomock.Controller
	recorder *MockIHasherMockRecorder
}

// MockIHasherMockRecorder is the mock recorder for MockIHasher.
type MockIHasherMockRecorder struct {
	mock *MockIHasher
}

// NewMockIHasher creates a new mock instance.
func NewMockIHasher(ctrl *gomock.Controller) *MockIHasher {
	mock := &MockIHasher{ctrl: ctrl}
	mock.recorder = &MockIHasherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHasher) EXPECT() *MockIHasherMockRecorder {
	return m.recorder
}

// CheckPassword mocks base method.
func (m *MockIHasher) CheckPassword(userPassword, providedPassword string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPassword", userPassword, providedPassword)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckPassword indicates an expected call of CheckPassword.
func (mr *MockIHasherMockRecorder) CheckPassword(userPassword, providedPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPassword", reflect.TypeOf((*MockIHasher)(nil).CheckPassword), userPassword, providedPassword)
}

// HashPassword mocks base method.
func (m *MockIHasher) HashPassword(password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashPassword", password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HashPassword indicates an expected call of HashPassword.
func (mr *MockIHasherMockRecorder) HashPassword(password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashPassword", reflect.TypeOf((*MockIHasher)(nil).HashPassword), password)
}
