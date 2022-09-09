package test

import (
	"reflect"

	"InnoTaxi/internal/app/auth"
	"InnoTaxi/internal/pkg/configs"

	"github.com/golang/mock/gomock"
)

// MockTokenForger is a mock of TokenForger interface.
type MockTokenForger struct {
	ctrl     *gomock.Controller
	recorder *MockTokenForgerMockRecorder
}

// MockTokenForgerMockRecorder is the mock recorder for MockTokenForger.
type MockTokenForgerMockRecorder struct {
	mock *MockTokenForger
}

// NewMockTokenForger creates a new mock instance.
func NewMockTokenForger(ctrl *gomock.Controller) *MockTokenForger {
	mock := &MockTokenForger{ctrl: ctrl}
	mock.recorder = &MockTokenForgerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenForger) EXPECT() *MockTokenForgerMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockTokenForger) Decode(cipher string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", cipher)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockTokenForgerMockRecorder) Decode(cipher interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockTokenForger)(nil).Decode), cipher)
}

// Encode mocks base method.
func (m *MockTokenForger) Encode(tokenClaims auth.TokenClaims, config configs.Config) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", tokenClaims, config)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockTokenForgerMockRecorder) Encode(tokenClaims, config interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockTokenForger)(nil).Encode), tokenClaims, config)
}
