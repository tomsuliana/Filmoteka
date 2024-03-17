package postgres

import (
	"database/sql"
	"server/server/internal/Film/repository"
	"server/server/internal/domain/entity"
)

type FilmRepo struct {
	DB *sql.DB
}

func NewFilmRepo(db *sql.DB) repository.FilmRepositoryI {
	return &FilmRepo{
		DB: db,
	}
}

func (repo *FilmRepo) CreateFilm(film *entity.Film) (uint, error) {
	insertFilm := `INSERT INTO film (name, description, release_date, rating) VALUES ($1, $2, $3, $4) RETURNING ID`
	var Id uint
	err := repo.DB.QueryRow(insertFilm, film.Name, film.Description, film.ReleaseDate, film.Rating).Scan(&Id)
	if err != nil {
		return 0, err
	}
	return Id, nil
}

func (repo *FilmRepo) AddActorToFilm(actorId uint, filmId uint) error {
	insertActorFilm := `INSERT INTO actor_film (actor_id, film_id) VALUES ($1, $2) RETURNING ID`
	var Id uint
	err := repo.DB.QueryRow(insertActorFilm, actorId, filmId).Scan(&Id)
	if err != nil {
		return err
	}
	return nil
}
