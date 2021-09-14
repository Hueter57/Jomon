// Code generated by MockGen. DO NOT EDIT.
// Source: model/transaction_tag.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	gomock "github.com/golang/mock/gomock"
)

// MockTransactionTagRepository is a mock of TransactionTagRepository interface
type MockTransactionTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionTagRepositoryMockRecorder
}

// MockTransactionTagRepositoryMockRecorder is the mock recorder for MockTransactionTagRepository
type MockTransactionTagRepositoryMockRecorder struct {
	mock *MockTransactionTagRepository
}

// NewMockTransactionTagRepository creates a new mock instance
func NewMockTransactionTagRepository(ctrl *gomock.Controller) *MockTransactionTagRepository {
	mock := &MockTransactionTagRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransactionTagRepository) EXPECT() *MockTransactionTagRepositoryMockRecorder {
	return m.recorder
}
