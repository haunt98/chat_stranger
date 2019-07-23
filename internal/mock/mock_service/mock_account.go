// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	model "github.com/1612180/chat_stranger/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccountService is a mock of AccountService interface
type MockAccountService struct {
	ctrl     *gomock.Controller
	recorder *MockAccountServiceMockRecorder
}

// MockAccountServiceMockRecorder is the mock recorder for MockAccountService
type MockAccountServiceMockRecorder struct {
	mock *MockAccountService
}

// NewMockAccountService creates a new mock instance
func NewMockAccountService(ctrl *gomock.Controller) *MockAccountService {
	mock := &MockAccountService{ctrl: ctrl}
	mock.recorder = &MockAccountServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccountService) EXPECT() *MockAccountServiceMockRecorder {
	return m.recorder
}

// SignUp mocks base method
func (m *MockAccountService) SignUp(showName, regName, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", showName, regName, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp
func (mr *MockAccountServiceMockRecorder) SignUp(showName, regName, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockAccountService)(nil).SignUp), showName, regName, password)
}

// LogIn mocks base method
func (m *MockAccountService) LogIn(regName, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogIn", regName, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LogIn indicates an expected call of LogIn
func (mr *MockAccountServiceMockRecorder) LogIn(regName, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogIn", reflect.TypeOf((*MockAccountService)(nil).LogIn), regName, password)
}

// Info mocks base method
func (m *MockAccountService) Info(userID int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Info", userID)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info
func (mr *MockAccountServiceMockRecorder) Info(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockAccountService)(nil).Info), userID)
}

// UpdateInfo mocks base method
func (m *MockAccountService) UpdateInfo(userID int, showName, gender string, birthYear int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInfo", userID, showName, gender, birthYear)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateInfo indicates an expected call of UpdateInfo
func (mr *MockAccountServiceMockRecorder) UpdateInfo(userID, showName, gender, birthYear interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInfo", reflect.TypeOf((*MockAccountService)(nil).UpdateInfo), userID, showName, gender, birthYear)
}
