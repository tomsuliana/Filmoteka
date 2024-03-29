// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	entity "server/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockFilmRepositoryI is a mock of FilmRepositoryI interface.
type MockFilmRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockFilmRepositoryIMockRecorder
}

// MockFilmRepositoryIMockRecorder is the mock recorder for MockFilmRepositoryI.
type MockFilmRepositoryIMockRecorder struct {
	mock *MockFilmRepositoryI
}

// NewMockFilmRepositoryI creates a new mock instance.
func NewMockFilmRepositoryI(ctrl *gomock.Controller) *MockFilmRepositoryI {
	mock := &MockFilmRepositoryI{ctrl: ctrl}
	mock.recorder = &MockFilmRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilmRepositoryI) EXPECT() *MockFilmRepositoryIMockRecorder {
	return m.recorder
}

// AddActorToFilm mocks base method.
func (m *MockFilmRepositoryI) AddActorToFilm(actorId, filmId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddActorToFilm", actorId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddActorToFilm indicates an expected call of AddActorToFilm.
func (mr *MockFilmRepositoryIMockRecorder) AddActorToFilm(actorId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddActorToFilm", reflect.TypeOf((*MockFilmRepositoryI)(nil).AddActorToFilm), actorId, filmId)
}

// CreateFilm mocks base method.
func (m *MockFilmRepositoryI) CreateFilm(film *entity.Film) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFilm", film)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFilm indicates an expected call of CreateFilm.
func (mr *MockFilmRepositoryIMockRecorder) CreateFilm(film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFilm", reflect.TypeOf((*MockFilmRepositoryI)(nil).CreateFilm), film)
}

// DeleteActorFromFilm mocks base method.
func (m *MockFilmRepositoryI) DeleteActorFromFilm(actorId, filmId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteActorFromFilm", actorId, filmId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteActorFromFilm indicates an expected call of DeleteActorFromFilm.
func (mr *MockFilmRepositoryIMockRecorder) DeleteActorFromFilm(actorId, filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteActorFromFilm", reflect.TypeOf((*MockFilmRepositoryI)(nil).DeleteActorFromFilm), actorId, filmId)
}

// DeleteFilm mocks base method.
func (m *MockFilmRepositoryI) DeleteFilm(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFilm", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFilm indicates an expected call of DeleteFilm.
func (mr *MockFilmRepositoryIMockRecorder) DeleteFilm(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilm", reflect.TypeOf((*MockFilmRepositoryI)(nil).DeleteFilm), id)
}

// GetActorsByFilm mocks base method.
func (m *MockFilmRepositoryI) GetActorsByFilm(filmId uint) ([]*entity.Actor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActorsByFilm", filmId)
	ret0, _ := ret[0].([]*entity.Actor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActorsByFilm indicates an expected call of GetActorsByFilm.
func (mr *MockFilmRepositoryIMockRecorder) GetActorsByFilm(filmId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActorsByFilm", reflect.TypeOf((*MockFilmRepositoryI)(nil).GetActorsByFilm), filmId)
}

// GetFilmById mocks base method.
func (m *MockFilmRepositoryI) GetFilmById(id uint) (*entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmById", id)
	ret0, _ := ret[0].(*entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmById indicates an expected call of GetFilmById.
func (mr *MockFilmRepositoryIMockRecorder) GetFilmById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmById", reflect.TypeOf((*MockFilmRepositoryI)(nil).GetFilmById), id)
}

// GetFilms mocks base method.
func (m *MockFilmRepositoryI) GetFilms(name, releaseDate bool) ([]*entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilms", name, releaseDate)
	ret0, _ := ret[0].([]*entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilms indicates an expected call of GetFilms.
func (mr *MockFilmRepositoryIMockRecorder) GetFilms(name, releaseDate interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilms", reflect.TypeOf((*MockFilmRepositoryI)(nil).GetFilms), name, releaseDate)
}

// GetFilmsByActor mocks base method.
func (m *MockFilmRepositoryI) GetFilmsByActor(actor *entity.Actor) ([]*entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsByActor", actor)
	ret0, _ := ret[0].([]*entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmsByActor indicates an expected call of GetFilmsByActor.
func (mr *MockFilmRepositoryIMockRecorder) GetFilmsByActor(actor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsByActor", reflect.TypeOf((*MockFilmRepositoryI)(nil).GetFilmsByActor), actor)
}

// SearchFilms mocks base method.
func (m *MockFilmRepositoryI) SearchFilms(word string) ([]*entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchFilms", word)
	ret0, _ := ret[0].([]*entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchFilms indicates an expected call of SearchFilms.
func (mr *MockFilmRepositoryIMockRecorder) SearchFilms(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchFilms", reflect.TypeOf((*MockFilmRepositoryI)(nil).SearchFilms), word)
}

// UpdateFilm mocks base method.
func (m *MockFilmRepositoryI) UpdateFilm(film *entity.Film) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilm", film)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFilm indicates an expected call of UpdateFilm.
func (mr *MockFilmRepositoryIMockRecorder) UpdateFilm(film interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilm", reflect.TypeOf((*MockFilmRepositoryI)(nil).UpdateFilm), film)
}
