// Code generated by MockGen. DO NOT EDIT.
// Source: message.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "github.com/1612180/chat_stranger/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockMessageRepo is a mock of MessageRepo interface
type MockMessageRepo struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepoMockRecorder
}

// MockMessageRepoMockRecorder is the mock recorder for MockMessageRepo
type MockMessageRepoMockRecorder struct {
	mock *MockMessageRepo
}

// NewMockMessageRepo creates a new mock instance
func NewMockMessageRepo(ctrl *gomock.Controller) *MockMessageRepo {
	mock := &MockMessageRepo{ctrl: ctrl}
	mock.recorder = &MockMessageRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageRepo) EXPECT() *MockMessageRepoMockRecorder {
	return m.recorder
}

// FetchByTime mocks base method
func (m *MockMessageRepo) FetchByTime(roomID int, fromTime time.Time) ([]*model.Message, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchByTime", roomID, fromTime)
	ret0, _ := ret[0].([]*model.Message)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FetchByTime indicates an expected call of FetchByTime
func (mr *MockMessageRepoMockRecorder) FetchByTime(roomID, fromTime interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchByTime", reflect.TypeOf((*MockMessageRepo)(nil).FetchByTime), roomID, fromTime)
}

// Create mocks base method
func (m *MockMessageRepo) Create(message *model.Message) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", message)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockMessageRepoMockRecorder) Create(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMessageRepo)(nil).Create), message)
}

// Delete mocks base method
func (m *MockMessageRepo) Delete(roomID int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", roomID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockMessageRepoMockRecorder) Delete(roomID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMessageRepo)(nil).Delete), roomID)
}
