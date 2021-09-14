// Code generated by MockGen. DO NOT EDIT.
// Source: model/transaction_detail.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	gomock "github.com/golang/mock/gomock"
)

// MockTransactionDetailRepository is a mock of TransactionDetailRepository interface
type MockTransactionDetailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionDetailRepositoryMockRecorder
}

// MockTransactionDetailRepositoryMockRecorder is the mock recorder for MockTransactionDetailRepository
type MockTransactionDetailRepositoryMockRecorder struct {
	mock *MockTransactionDetailRepository
}

// NewMockTransactionDetailRepository creates a new mock instance
func NewMockTransactionDetailRepository(ctrl *gomock.Controller) *MockTransactionDetailRepository {
	mock := &MockTransactionDetailRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionDetailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionDetailRepository) EXPECT() *MockTransactionDetailRepositoryMockRecorder {
	return m.recorder
}
