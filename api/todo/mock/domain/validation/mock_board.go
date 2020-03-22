// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/validation/board.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	context "context"
	domain "github.com/16francs/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBoardDomainValidation is a mock of BoardDomainValidation interface
type MockBoardDomainValidation struct {
	ctrl     *gomock.Controller
	recorder *MockBoardDomainValidationMockRecorder
}

// MockBoardDomainValidationMockRecorder is the mock recorder for MockBoardDomainValidation
type MockBoardDomainValidationMockRecorder struct {
	mock *MockBoardDomainValidation
}

// NewMockBoardDomainValidation creates a new mock instance
func NewMockBoardDomainValidation(ctrl *gomock.Controller) *MockBoardDomainValidation {
	mock := &MockBoardDomainValidation{ctrl: ctrl}
	mock.recorder = &MockBoardDomainValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBoardDomainValidation) EXPECT() *MockBoardDomainValidationMockRecorder {
	return m.recorder
}

// Board mocks base method
func (m *MockBoardDomainValidation) Board(ctx context.Context, b *domain.Board) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Board", ctx, b)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// Board indicates an expected call of Board
func (mr *MockBoardDomainValidationMockRecorder) Board(ctx, b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Board", reflect.TypeOf((*MockBoardDomainValidation)(nil).Board), ctx, b)
}

// BoardList mocks base method
func (m *MockBoardDomainValidation) BoardList(ctx context.Context, bl *domain.BoardList) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BoardList", ctx, bl)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// BoardList indicates an expected call of BoardList
func (mr *MockBoardDomainValidationMockRecorder) BoardList(ctx, bl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BoardList", reflect.TypeOf((*MockBoardDomainValidation)(nil).BoardList), ctx, bl)
}
