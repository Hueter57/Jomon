// Code generated by MockGen. DO NOT EDIT.
// Source: service/service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	gomock "github.com/golang/mock/gomock"
	service "github.com/traPtitech/Jomon/service"
	io "io"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CreateFile mocks base method
func (m *MockService) CreateFile(src io.Reader, mimetype string) (service.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFile", src, mimetype)
	ret0, _ := ret[0].(service.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFile indicates an expected call of CreateFile
func (mr *MockServiceMockRecorder) CreateFile(src, mimetype interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFile", reflect.TypeOf((*MockService)(nil).CreateFile), src, mimetype)
}

// GetAccessToken mocks base method
func (m *MockService) GetAccessToken(code, codeVerifier string) (service.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccessToken", code, codeVerifier)
	ret0, _ := ret[0].(service.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccessToken indicates an expected call of GetAccessToken
func (mr *MockServiceMockRecorder) GetAccessToken(code, codeVerifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccessToken", reflect.TypeOf((*MockService)(nil).GetAccessToken), code, codeVerifier)
}

// GetClientId mocks base method
func (m *MockService) GetClientId() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientId")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetClientId indicates an expected call of GetClientId
func (mr *MockServiceMockRecorder) GetClientId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientId", reflect.TypeOf((*MockService)(nil).GetClientId))
}
