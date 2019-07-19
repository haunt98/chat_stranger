// Code generated by MockGen. DO NOT EDIT.
// Source: room.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	model "github.com/1612180/chat_stranger/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRoomRepository is a mock of RoomRepository interface
type MockRoomRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRoomRepositoryMockRecorder
}

// MockRoomRepositoryMockRecorder is the mock recorder for MockRoomRepository
type MockRoomRepositoryMockRecorder struct {
	mock *MockRoomRepository
}

// NewMockRoomRepository creates a new mock instance
func NewMockRoomRepository(ctrl *gomock.Controller) *MockRoomRepository {
	mock := &MockRoomRepository{ctrl: ctrl}
	mock.recorder = &MockRoomRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRoomRepository) EXPECT() *MockRoomRepositoryMockRecorder {
	return m.recorder
}

// Exist mocks base method
func (m *MockRoomRepository) Exist(id int) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exist", id)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exist indicates an expected call of Exist
func (mr *MockRoomRepositoryMockRecorder) Exist(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exist", reflect.TypeOf((*MockRoomRepository)(nil).Exist), id)
}

// FindEmpty mocks base method
func (m *MockRoomRepository) FindEmpty() (*model.Room, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindEmpty")
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindEmpty indicates an expected call of FindEmpty
func (mr *MockRoomRepositoryMockRecorder) FindEmpty() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindEmpty", reflect.TypeOf((*MockRoomRepository)(nil).FindEmpty))
}

// FindNext mocks base method
func (m *MockRoomRepository) FindNext(old int) (*model.Room, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindNext", old)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindNext indicates an expected call of FindNext
func (mr *MockRoomRepositoryMockRecorder) FindNext(old interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindNext", reflect.TypeOf((*MockRoomRepository)(nil).FindNext), old)
}

// FindSameGender mocks base method
func (m *MockRoomRepository) FindSameGender(old int, gender string) (*model.Room, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSameGender", old, gender)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindSameGender indicates an expected call of FindSameGender
func (mr *MockRoomRepositoryMockRecorder) FindSameGender(old, gender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSameGender", reflect.TypeOf((*MockRoomRepository)(nil).FindSameGender), old, gender)
}

// FindSameBirthYear mocks base method
func (m *MockRoomRepository) FindSameBirthYear(old, year int) (*model.Room, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindSameBirthYear", old, year)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindSameBirthYear indicates an expected call of FindSameBirthYear
func (mr *MockRoomRepositoryMockRecorder) FindSameBirthYear(old, year interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSameBirthYear", reflect.TypeOf((*MockRoomRepository)(nil).FindSameBirthYear), old, year)
}

// FindByUser mocks base method
func (m *MockRoomRepository) FindByUser(userID int) (*model.Room, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUser", userID)
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// FindByUser indicates an expected call of FindByUser
func (mr *MockRoomRepositoryMockRecorder) FindByUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUser", reflect.TypeOf((*MockRoomRepository)(nil).FindByUser), userID)
}

// Create mocks base method
func (m *MockRoomRepository) Create() (*model.Room, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(*model.Room)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockRoomRepositoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRoomRepository)(nil).Create))
}
