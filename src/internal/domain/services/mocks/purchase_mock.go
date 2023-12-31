// Code generated by MockGen. DO NOT EDIT.
// Source: sd/internal/domain/services (interfaces: IPurchaseRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	entities "sd/internal/domain/entities"

	gomock "github.com/golang/mock/gomock"
)

// MockIPurchaseRepo is a mock of IPurchaseRepo interface.
type MockIPurchaseRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIPurchaseRepoMockRecorder
}

// MockIPurchaseRepoMockRecorder is the mock recorder for MockIPurchaseRepo.
type MockIPurchaseRepoMockRecorder struct {
	mock *MockIPurchaseRepo
}

// NewMockIPurchaseRepo creates a new mock instance.
func NewMockIPurchaseRepo(ctrl *gomock.Controller) *MockIPurchaseRepo {
	mock := &MockIPurchaseRepo{ctrl: ctrl}
	mock.recorder = &MockIPurchaseRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPurchaseRepo) EXPECT() *MockIPurchaseRepoMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIPurchaseRepo) Create(arg0 context.Context, arg1 *entities.Purchase) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIPurchaseRepoMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIPurchaseRepo)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockIPurchaseRepo) Delete(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockIPurchaseRepoMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIPurchaseRepo)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockIPurchaseRepo) GetAll(arg0 context.Context) (entities.Purchases, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0)
	ret0, _ := ret[0].(entities.Purchases)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockIPurchaseRepoMockRecorder) GetAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockIPurchaseRepo)(nil).GetAll), arg0)
}

// GetById mocks base method.
func (m *MockIPurchaseRepo) GetById(arg0 context.Context, arg1 int) (*entities.Purchase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*entities.Purchase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockIPurchaseRepoMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockIPurchaseRepo)(nil).GetById), arg0, arg1)
}

// Update mocks base method.
func (m *MockIPurchaseRepo) Update(arg0 context.Context, arg1 *entities.Purchase) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockIPurchaseRepoMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIPurchaseRepo)(nil).Update), arg0, arg1)
}
