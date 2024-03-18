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

func (repo *FilmRepo) UpdateFilm(film *entity.Film) error {
	updateFilm := `UPDATE film 
				   SET name = $1, description = $2, release_date = $3, rating = $4
				   WHERE id = $5`
	_, err := repo.DB.Exec(updateFilm, film.Name, film.Description, film.ReleaseDate, film.Rating, film.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FilmRepo) GetFilmById(id uint) (*entity.Film, error) {
	film := &entity.Film{}
	row := repo.DB.QueryRow(`SELECT id, name, description, release_date, rating FROM film WHERE id = $1`, id)
	err := row.Scan(&film.ID,
		&film.Name,
		&film.Description,
		&film.ReleaseDate,
		&film.Rating,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return film, nil
}

func (repo *FilmRepo) DeleteActorFromFilm(actorId uint, filmId uint) error {
	deleteActor := `DELETE FROM actor_film WHERE actor_id = $1 AND film_id = $2`
	_, err := repo.DB.Exec(deleteActor, actorId, filmId)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FilmRepo) DeleteFilm(id uint) error {
	deleteActors := `DELETE FROM actor_film WHERE film_id = $1`
	_, err := repo.DB.Exec(deleteActors, id)
	if err != nil {
		return err
	}

	deletFilm := `DELETE FROM film WHERE id = $1`
	_, err = repo.DB.Exec(deletFilm, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *FilmRepo) GetFilms(name bool, releaseDate bool) ([]*entity.Film, error) {
	filmsQuery := `SELECT id, name, description, release_date, rating
													FROM film 
													ORDER BY rating DESC`

	if name {
		filmsQuery = filmsQuery + `, name ASC`
	}
	if releaseDate {
		filmsQuery = filmsQuery + `, release_date DESC`
	}
	rows, err := repo.DB.Query(filmsQuery)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Films = []*entity.Film{}
	for rows.Next() {
		film := &entity.Film{}
		err = rows.Scan(
			&film.ID,
			&film.Name,
			&film.Description,
			&film.ReleaseDate,
			&film.Rating,
		)
		if err != nil {
			return nil, err
		}
		Films = append(Films, film)
	}
	return Films, nil
}

func (repo *FilmRepo) GetActorsByFilm(filmId uint) ([]*entity.Actor, error) {
	rows, err := repo.DB.Query(`SELECT actor.id, actor.name, surname, birthday, gender 
								FROM actor_film af
								INNER JOIN actor ON af.actor_id=actor.id 
								WHERE film_id = $1`, filmId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Actors = []*entity.Actor{}
	var count = 0
	for rows.Next() {
		count++
		actor := &entity.Actor{}
		err = rows.Scan(
			&actor.ID,
			&actor.Name,
			&actor.Surname,
			&actor.Birthday,
			&actor.Gender,
		)
		if err != nil {
			return nil, err
		}
		Actors = append(Actors, actor)
	}
	return Actors, nil
}

func (repo *FilmRepo) SearchFilms(word string) ([]*entity.Film, error) {
	rows, err := repo.DB.Query(`SELECT id, name, description, release_date, rating
							    FROM film 
								WHERE LOWER(name) 
								LIKE LOWER('%' || $1 || '%')
								ORDER BY rating DESC`, word)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Films = []*entity.Film{}
	for rows.Next() {
		film := &entity.Film{}
		err = rows.Scan(
			&film.ID,
			&film.Name,
			&film.Description,
			&film.ReleaseDate,
			&film.Rating,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		Films = append(Films, film)
	}
	return Films, nil
}

func (repo *FilmRepo) GetFilmsByActor(actor *entity.Actor) ([]*entity.Film, error) {
	rows, err := repo.DB.Query(`SELECT film.id, film.name, description, release_date, rating
								FROM actor_film af 
								INNER JOIN film ON af.film_id=film.id
								INNER JOIN actor ON af.actor_id=actor.id 
								WHERE actor.id = $1
								ORDER BY rating DESC`, actor.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	defer rows.Close()
	var Films = []*entity.Film{}
	var count = 0
	for rows.Next() {
		count++
		film := &entity.Film{}
		err = rows.Scan(
			&film.ID,
			&film.Name,
			&film.Description,
			&film.ReleaseDate,
			&film.Rating,
		)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		Films = append(Films, film)
	}
	if count == 0 {
		return nil, entity.ErrNotFound
	}
	return Films, nil
}
