package test

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockICacheRepository is a mock of ICacheRepository interface.
type MockICacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICacheRepositoryMockRecorder
}

// MockICacheRepositoryMockRecorder is the mock recorder for MockICacheRepository.
type MockICacheRepositoryMockRecorder struct {
	mock *MockICacheRepository
}

// NewMockICacheRepository creates a new mock instance.
func NewMockICacheRepository(ctrl *gomock.Controller) *MockICacheRepository {
	mock := &MockICacheRepository{ctrl: ctrl}
	mock.recorder = &MockICacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICacheRepository) EXPECT() *MockICacheRepositoryMockRecorder {
	return m.recorder
}

// CreateInvalidToken mocks base method.
func (m *MockICacheRepository) CreateInvalidToken(ctx context.Context, token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvalidToken", ctx, token)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInvalidToken indicates an expected call of CreateInvalidToken.
func (mr *MockICacheRepositoryMockRecorder) CreateInvalidToken(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvalidToken", reflect.TypeOf((*MockICacheRepository)(nil).CreateInvalidToken), ctx, token)
}

// DoesInvalidTokenExist mocks base method.
func (m *MockICacheRepository) DoesInvalidTokenExist(ctx context.Context, token string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoesInvalidTokenExist", ctx, token)
	ret0, _ := ret[0].(bool)
	return ret0
}

// DoesInvalidTokenExist indicates an expected call of DoesInvalidTokenExist.
func (mr *MockICacheRepositoryMockRecorder) DoesInvalidTokenExist(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoesInvalidTokenExist", reflect.TypeOf((*MockICacheRepository)(nil).DoesInvalidTokenExist), ctx, token)
}
