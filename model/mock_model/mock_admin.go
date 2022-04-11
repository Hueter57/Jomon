// Code generated by MockGen. DO NOT EDIT.
// Source: admin.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	model "github.com/traPtitech/Jomon/model"
	reflect "reflect"
)

// MockAdminRepository is a mock of AdminRepository interface
type MockAdminRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAdminRepositoryMockRecorder
}

// MockAdminRepositoryMockRecorder is the mock recorder for MockAdminRepository
type MockAdminRepositoryMockRecorder struct {
	mock *MockAdminRepository
}

// NewMockAdminRepository creates a new mock instance
func NewMockAdminRepository(ctrl *gomock.Controller) *MockAdminRepository {
	mock := &MockAdminRepository{ctrl: ctrl}
	mock.recorder = &MockAdminRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAdminRepository) EXPECT() *MockAdminRepositoryMockRecorder {
	return m.recorder
}

// GetAdmins mocks base method
func (m *MockAdminRepository) GetAdmins(ctx context.Context) ([]*model.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmins", ctx)
	ret0, _ := ret[0].([]*model.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmins indicates an expected call of GetAdmins
func (mr *MockAdminRepositoryMockRecorder) GetAdmins(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmins", reflect.TypeOf((*MockAdminRepository)(nil).GetAdmins), ctx)
}

// AddAdmins mocks base method
func (m *MockAdminRepository) AddAdmins(ctx context.Context, userIDs []uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAdmins", ctx, userIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddAdmins indicates an expected call of AddAdmins
func (mr *MockAdminRepositoryMockRecorder) AddAdmins(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAdmins", reflect.TypeOf((*MockAdminRepository)(nil).AddAdmins), ctx, userIDs)
}

// DeleteAdmins mocks base method
func (m *MockAdminRepository) DeleteAdmins(ctx context.Context, userIDs []uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmins", ctx, userIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdmins indicates an expected call of DeleteAdmins
func (mr *MockAdminRepositoryMockRecorder) DeleteAdmins(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmins", reflect.TypeOf((*MockAdminRepository)(nil).DeleteAdmins), ctx, userIDs)
}
