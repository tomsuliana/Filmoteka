// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	entity "server/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockActorRepositoryI is a mock of ActorRepositoryI interface.
type MockActorRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockActorRepositoryIMockRecorder
}

// MockActorRepositoryIMockRecorder is the mock recorder for MockActorRepositoryI.
type MockActorRepositoryIMockRecorder struct {
	mock *MockActorRepositoryI
}

// NewMockActorRepositoryI creates a new mock instance.
func NewMockActorRepositoryI(ctrl *gomock.Controller) *MockActorRepositoryI {
	mock := &MockActorRepositoryI{ctrl: ctrl}
	mock.recorder = &MockActorRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockActorRepositoryI) EXPECT() *MockActorRepositoryIMockRecorder {
	return m.recorder
}

// CreateActor mocks base method.
func (m *MockActorRepositoryI) CreateActor(actor *entity.Actor) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateActor", actor)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateActor indicates an expected call of CreateActor.
func (mr *MockActorRepositoryIMockRecorder) CreateActor(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateActor", reflect.TypeOf((*MockActorRepositoryI)(nil).CreateActor), actor)
}

// DeleteActor mocks base method.
func (m *MockActorRepositoryI) DeleteActor(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActor", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActor indicates an expected call of DeleteActor.
func (mr *MockActorRepositoryIMockRecorder) DeleteActor(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActor", reflect.TypeOf((*MockActorRepositoryI)(nil).DeleteActor), id)
}

// GetActorById mocks base method.
func (m *MockActorRepositoryI) GetActorById(id uint) (*entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorById", id)
	ret0, _ := ret[0].(*entity.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorById indicates an expected call of GetActorById.
func (mr *MockActorRepositoryIMockRecorder) GetActorById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorById", reflect.TypeOf((*MockActorRepositoryI)(nil).GetActorById), id)
}

// GetActorByName mocks base method.
func (m *MockActorRepositoryI) GetActorByName(name, surname string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorByName", name, surname)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorByName indicates an expected call of GetActorByName.
func (mr *MockActorRepositoryIMockRecorder) GetActorByName(name, surname interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorByName", reflect.TypeOf((*MockActorRepositoryI)(nil).GetActorByName), name, surname)
}

// GetActors mocks base method.
func (m *MockActorRepositoryI) GetActors() ([]*entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActors")
	ret0, _ := ret[0].([]*entity.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActors indicates an expected call of GetActors.
func (mr *MockActorRepositoryIMockRecorder) GetActors() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActors", reflect.TypeOf((*MockActorRepositoryI)(nil).GetActors))
}

// SearchActors mocks base method.
func (m *MockActorRepositoryI) SearchActors(word string) ([]*entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchActors", word)
	ret0, _ := ret[0].([]*entity.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchActors indicates an expected call of SearchActors.
func (mr *MockActorRepositoryIMockRecorder) SearchActors(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchActors", reflect.TypeOf((*MockActorRepositoryI)(nil).SearchActors), word)
}

// UpdateActor mocks base method.
func (m *MockActorRepositoryI) UpdateActor(actor *entity.Actor) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateActor", actor)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateActor indicates an expected call of UpdateActor.
func (mr *MockActorRepositoryIMockRecorder) UpdateActor(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateActor", reflect.TypeOf((*MockActorRepositoryI)(nil).UpdateActor), actor)
}
