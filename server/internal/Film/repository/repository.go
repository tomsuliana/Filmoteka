package repository

import (
	"server/server/internal/domain/entity"
)

type FilmRepositoryI interface {
	CreateFilm(film *entity.Film) (uint, error)
	AddActorToFilm(actorId uint, filmId uint) error
	UpdateFilm(film *entity.Film) error
	GetFilmById(id uint) (*entity.Film, error)
	DeleteActorFromFilm(actorId uint, filmId uint) error
	DeleteFilm(id uint) error
	GetFilms(name bool, releaseDate bool) ([]*entity.Film, error)
	GetActorsByFilm(filmId uint) ([]*entity.Actor, error)
}
