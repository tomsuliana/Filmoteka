package postgres

import (
	// "errors"
	"reflect"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateFilmSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	film := &entity.Film{
		Name:        "Terminator 2",
		Description: "Cool film",
		ReleaseDate: "1985-08-22",
		Rating:      5.9,
	}

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectQuery(`INSERT INTO film`).
		WithArgs(film.Name, film.Description, film.ReleaseDate, film.Rating).
		WillReturnRows(rows)

	repo := &FilmRepo{
		DB: db,
	}

	id, err := repo.CreateFilm(film)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if id != 1 {
		t.Errorf("bad id: want %v, have %v", id, 1)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddActorToFilmSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var actorId, filmId uint
	actorId = 1
	filmId = 1

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectQuery(`INSERT INTO actor_film`).
		WithArgs(actorId, filmId).
		WillReturnRows(rows)

	repo := &FilmRepo{
		DB: db,
	}

	err = repo.AddActorToFilm(actorId, filmId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateFilmSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	film := &entity.Film{
		ID:          1,
		Description: "Coolest film",
	}

	mock.
		ExpectExec(`UPDATE film SET`).
		WithArgs(film.Name, film.Description, film.ReleaseDate, film.Rating, film.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &FilmRepo{
		DB: db,
	}

	err = repo.UpdateFilm(film)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFilmByIdSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var filmId uint
	filmId = 1

	film := &entity.Film{
		ID:          1,
		Name:        "Terminator 2",
		Description: "Cool film",
		ReleaseDate: "1985-08-22",
		Rating:      5.9,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "description", "release_date", "rating"}).
		AddRow(film.ID, film.Name, film.Description, film.ReleaseDate, film.Rating)

	mock.
		ExpectQuery(`SELECT id, name, description, release_date, rating FROM film `).
		WithArgs(filmId).
		WillReturnRows(rows)

	repo := &FilmRepo{
		DB: db,
	}

	actual, err := repo.GetFilmById(filmId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(actual, film) {
		t.Errorf("results not match, want %v, have %v", film, actual)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteActorFromFilmSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var actorId, filmId uint
	actorId = 1
	filmId = 1

	mock.
		ExpectExec(`DELETE FROM actor_film WHERE`).
		WithArgs(actorId, filmId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &FilmRepo{
		DB: db,
	}

	err = repo.DeleteActorFromFilm(actorId, filmId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteFilmSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var filmId uint
	filmId = 1

	mock.
		ExpectExec(`DELETE FROM actor_film WHERE`).
		WithArgs(filmId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.
		ExpectExec(`DELETE FROM film WHERE`).
		WithArgs(filmId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &FilmRepo{
		DB: db,
	}

	err = repo.DeleteFilm(filmId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFilmsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &FilmRepo{
		DB: db,
	}

	name := true
	release_date := true

	rows := sqlmock.
		NewRows([]string{"id", "name", "description", "release_date", "rating"})
	expect := []*entity.Film{
		{
			ID:          1,
			Name:        "Terminator 2",
			Description: "Cool film",
			ReleaseDate: "1985-08-22",
			Rating:      5.9,
		},
		{
			ID:          2,
			Name:        "Terminator 3",
			Description: "Cool film",
			ReleaseDate: "1986-08-22",
			Rating:      6.9,
		},
	}
	for _, film := range expect {
		rows = rows.AddRow(film.ID, film.Name, film.Description, film.ReleaseDate, film.Rating)
	}

	mock.
		ExpectQuery(`SELECT id, name, description, release_date, rating
		FROM film `).
		WillReturnRows(rows)

	films, err := repo.GetFilms(name, release_date)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(films[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], films[0])
		return
	}
}

func TestGetActorsByFilmSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	filmid := 1

	repo := &FilmRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "surname", "birthday", "gender"})
	expect := []*entity.Actor{
		{
			ID:       1,
			Name:     "john",
			Surname:  "doe",
			Birthday: "1985-08-22",
			Gender:   "ж",
		},
		{
			ID:       2,
			Name:     "john",
			Surname:  "doe",
			Birthday: "1985-08-22",
			Gender:   "ж",
		},
	}
	for _, actor := range expect {
		rows = rows.AddRow(actor.ID, actor.Name, actor.Surname, actor.Birthday, actor.Gender)
	}

	mock.
		ExpectQuery(`SELECT actor.id, actor.name, surname, birthday, gender 
		FROM actor_film af
		INNER JOIN actor ON af.actor_id=actor.id `).
		WillReturnRows(rows)

	actors, err := repo.GetActorsByFilm(uint(filmid))
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(actors[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], actors[0])
		return
	}
}

func TestSearchFilmsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &FilmRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "description", "release_date", "rating"})
	expect := []*entity.Film{
		{
			ID:          1,
			Name:        "Terminator 2",
			Description: "Cool film",
			ReleaseDate: "1985-08-22",
			Rating:      5.9,
		},
		{
			ID:          2,
			Name:        "Terminator 3",
			Description: "Cool film",
			ReleaseDate: "1986-08-22",
			Rating:      6.9,
		},
	}
	for _, film := range expect {
		rows = rows.AddRow(film.ID, film.Name, film.Description, film.ReleaseDate, film.Rating)
	}

	mock.
		ExpectQuery(`SELECT id, name, description, release_date, rating
		FROM film 
		WHERE`).
		WillReturnRows(rows)

	name := "Terminator"

	films, err := repo.SearchFilms(name)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(films[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], films[0])
		return
	}
}

func TestGetFilmsByActorSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &FilmRepo{
		DB: db,
	}

	rows := sqlmock.
		NewRows([]string{"id", "name", "description", "release_date", "rating"})
	expect := []*entity.Film{
		{
			ID:          1,
			Name:        "Terminator 2",
			Description: "Cool film",
			ReleaseDate: "1985-08-22",
			Rating:      5.9,
		},
		{
			ID:          2,
			Name:        "Terminator 3",
			Description: "Cool film",
			ReleaseDate: "1986-08-22",
			Rating:      6.9,
		},
	}
	for _, film := range expect {
		rows = rows.AddRow(film.ID, film.Name, film.Description, film.ReleaseDate, film.Rating)
	}

	mock.
		ExpectQuery(`SELECT film.id, film.name, description, release_date, rating
		FROM actor_film af 
		INNER JOIN film ON af.film_id=film.id
		INNER JOIN actor ON af.actor_id=actor.id`).
		WillReturnRows(rows)

	actor := &entity.Actor{
		ID:       1,
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	films, err := repo.GetFilmsByActor(actor)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(films[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], films[0])
		return
	}
}
