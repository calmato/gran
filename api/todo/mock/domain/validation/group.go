// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/validation/group.go

// Package mock_validation is a generated GoMock package.
package mock_validation

import (
	context "context"
	domain "github.com/calmato/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGroupDomainValidation is a mock of GroupDomainValidation interface
type MockGroupDomainValidation struct {
	ctrl     *gomock.Controller
	recorder *MockGroupDomainValidationMockRecorder
}

// MockGroupDomainValidationMockRecorder is the mock recorder for MockGroupDomainValidation
type MockGroupDomainValidationMockRecorder struct {
	mock *MockGroupDomainValidation
}

// NewMockGroupDomainValidation creates a new mock instance
func NewMockGroupDomainValidation(ctrl *gomock.Controller) *MockGroupDomainValidation {
	mock := &MockGroupDomainValidation{ctrl: ctrl}
	mock.recorder = &MockGroupDomainValidationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGroupDomainValidation) EXPECT() *MockGroupDomainValidationMockRecorder {
	return m.recorder
}

// Group mocks base method
func (m *MockGroupDomainValidation) Group(ctx context.Context, g *domain.Group) []*domain.ValidationError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", ctx, g)
	ret0, _ := ret[0].([]*domain.ValidationError)
	return ret0
}

// Group indicates an expected call of Group
func (mr *MockGroupDomainValidationMockRecorder) Group(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockGroupDomainValidation)(nil).Group), ctx, g)
}
