package repository

import (
	"server/server/internal/domain/entity"
)

type FilmRepositoryI interface {
	CreateFilm(film *entity.Film) (uint, error)
	AddActorToFilm(actorId uint, filmId uint) error
}
