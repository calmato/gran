// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/service/group.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	domain "github.com/16francs/gran/api/todo/internal/domain"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGroupService is a mock of GroupService interface
type MockGroupService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupServiceMockRecorder
}

// MockGroupServiceMockRecorder is the mock recorder for MockGroupService
type MockGroupServiceMockRecorder struct {
	mock *MockGroupService
}

// NewMockGroupService creates a new mock instance
func NewMockGroupService(ctrl *gomock.Controller) *MockGroupService {
	mock := &MockGroupService{ctrl: ctrl}
	mock.recorder = &MockGroupServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGroupService) EXPECT() *MockGroupServiceMockRecorder {
	return m.recorder
}

// Index mocks base method
func (m *MockGroupService) Index(ctx context.Context, u *domain.User) ([]*domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Index", ctx, u)
	ret0, _ := ret[0].([]*domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Index indicates an expected call of Index
func (mr *MockGroupServiceMockRecorder) Index(ctx, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Index", reflect.TypeOf((*MockGroupService)(nil).Index), ctx, u)
}

// Show mocks base method
func (m *MockGroupService) Show(ctx context.Context, groupID string) (*domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Show", ctx, groupID)
	ret0, _ := ret[0].(*domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Show indicates an expected call of Show
func (mr *MockGroupServiceMockRecorder) Show(ctx, groupID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Show", reflect.TypeOf((*MockGroupService)(nil).Show), ctx, groupID)
}

// Create mocks base method
func (m *MockGroupService) Create(ctx context.Context, u *domain.User, g *domain.Group) (*domain.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, u, g)
	ret0, _ := ret[0].(*domain.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockGroupServiceMockRecorder) Create(ctx, u, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockGroupService)(nil).Create), ctx, u, g)
}

// Update mocks base method
func (m *MockGroupService) Update(ctx context.Context, g *domain.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockGroupServiceMockRecorder) Update(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockGroupService)(nil).Update), ctx, g)
}

// InviteUsers mocks base method
func (m *MockGroupService) InviteUsers(ctx context.Context, g *domain.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InviteUsers", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// InviteUsers indicates an expected call of InviteUsers
func (mr *MockGroupServiceMockRecorder) InviteUsers(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InviteUsers", reflect.TypeOf((*MockGroupService)(nil).InviteUsers), ctx, g)
}

// Join mocks base method
func (m *MockGroupService) Join(ctx context.Context, g *domain.Group) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Join", ctx, g)
	ret0, _ := ret[0].(error)
	return ret0
}

// Join indicates an expected call of Join
func (mr *MockGroupServiceMockRecorder) Join(ctx, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockGroupService)(nil).Join), ctx, g)
}

// IsContainInUserIDs mocks base method
func (m *MockGroupService) IsContainInUserIDs(ctx context.Context, userID string, g *domain.Group) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsContainInUserIDs", ctx, userID, g)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsContainInUserIDs indicates an expected call of IsContainInUserIDs
func (mr *MockGroupServiceMockRecorder) IsContainInUserIDs(ctx, userID, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsContainInUserIDs", reflect.TypeOf((*MockGroupService)(nil).IsContainInUserIDs), ctx, userID, g)
}

// IsContainInInvitedEmails mocks base method
func (m *MockGroupService) IsContainInInvitedEmails(ctx context.Context, email string, g *domain.Group) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsContainInInvitedEmails", ctx, email, g)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsContainInInvitedEmails indicates an expected call of IsContainInInvitedEmails
func (mr *MockGroupServiceMockRecorder) IsContainInInvitedEmails(ctx, email, g interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsContainInInvitedEmails", reflect.TypeOf((*MockGroupService)(nil).IsContainInInvitedEmails), ctx, email, g)
}
