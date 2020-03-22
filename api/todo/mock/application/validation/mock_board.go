// Code generated by MockGen. DO NOT EDIT.
// Source: internal/application/validation/board.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	request "github.com/16francs/gran/api/todo/internal/application/request"
	domain "github.com/16francs/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBoardRequestValidation is a mock of BoardRequestValidation interface
type MockBoardRequestValidation struct {
	ctrl     *gomock.Controller
	recorder *MockBoardRequestValidationMockRecorder
}

// MockBoardRequestValidationMockRecorder is the mock recorder for MockBoardRequestValidation
type MockBoardRequestValidationMockRecorder struct {
	mock *MockBoardRequestValidation
}

// NewMockBoardRequestValidation creates a new mock instance
func NewMockBoardRequestValidation(ctrl *gomock.Controller) *MockBoardRequestValidation {
	mock := &MockBoardRequestValidation{ctrl: ctrl}
	mock.recorder = &MockBoardRequestValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBoardRequestValidation) EXPECT() *MockBoardRequestValidationMockRecorder {
	return m.recorder
}

// CreateBoard mocks base method
func (m *MockBoardRequestValidation) CreateBoard(req *request.CreateBoard) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBoard", req)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// CreateBoard indicates an expected call of CreateBoard
func (mr *MockBoardRequestValidationMockRecorder) CreateBoard(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBoard", reflect.TypeOf((*MockBoardRequestValidation)(nil).CreateBoard), req)
}

// CreateBoardList mocks base method
func (m *MockBoardRequestValidation) CreateBoardList(req *request.CreateBoardList) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBoardList", req)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// CreateBoardList indicates an expected call of CreateBoardList
func (mr *MockBoardRequestValidationMockRecorder) CreateBoardList(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBoardList", reflect.TypeOf((*MockBoardRequestValidation)(nil).CreateBoardList), req)
}
