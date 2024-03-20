package postgres

import (
	"errors"
	"reflect"
	"server/internal/domain/entity"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestCreateActorSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	actor := &entity.Actor{
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	rows := sqlmock.
		NewRows([]string{"id"}).AddRow(1)

	mock.
		ExpectQuery(`INSERT INTO actor`).
		WithArgs(actor.Name, actor.Surname, actor.Birthday, actor.Gender).
		WillReturnRows(rows)

	repo := &ActorRepo{
		DB: db,
	}

	id, err := repo.CreateActor(actor)
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

func TestCreateActorFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	actor := &entity.Actor{
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	testErr := errors.New("testErr")

	mock.
		ExpectQuery("INSERT INTO actor").
		WithArgs(actor.Name, actor.Surname, actor.Birthday, actor.Gender).
		WillReturnError(testErr)

	repo := &ActorRepo{
		DB: db,
	}

	_, err = repo.CreateActor(actor)
	if err != testErr {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestUpdateActorSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	actor := &entity.Actor{
		ID:       1,
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	mock.
		ExpectExec(`UPDATE actor SET `).
		WithArgs(actor.Name, actor.Surname, actor.Birthday, actor.Gender, actor.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo := &ActorRepo{
		DB: db,
	}

	err = repo.UpdateActor(actor)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetActorById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var id uint = 2

	rows := sqlmock.
		NewRows([]string{"id", "name", "surname", "birthday", "gender"})

	expect := &entity.Actor{
		ID:       1,
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	rows = rows.AddRow(expect.ID, expect.Name, expect.Surname, expect.Birthday, expect.Gender)

	mock.
		ExpectQuery("SELECT id, name, surname, birthday, gender FROM actor WHERE").
		WithArgs(id).
		WillReturnRows(rows)

	repo := &ActorRepo{
		DB: db,
	}

	user, err := repo.GetActorById(id)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
	if !reflect.DeepEqual(user, expect) {
		t.Errorf("results not match, want %v, have %v", expect, user)
		return
	}
}

func TestDeleteActorSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &ActorRepo{
		DB: db,
	}

	var actorID uint
	actorID = 1
	mock.
		ExpectExec(`DELETE FROM actor WHERE`).
		WithArgs(actorID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteActor(actorID)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestGetActorByNameSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	var name string = "john"
	var surname string = "doe"

	rows := sqlmock.
		NewRows([]string{"id"})

	expect := 1

	rows = rows.AddRow(expect)

	mock.
		ExpectQuery("SELECT id FROM actor WHERE").
		WithArgs(name, surname).
		WillReturnRows(rows)

	repo := &ActorRepo{
		DB: db,
	}

	_, err = repo.GetActorByName(name, surname)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
		return
	}
}

func TestSearchActorsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &ActorRepo{
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

	var word = "Ani"

	mock.
		ExpectQuery(`SELECT id, name, surname, birthday, gender
		FROM actor 
		WHERE `).WithArgs(word).
		WillReturnRows(rows)

	actors, err := repo.SearchActors(word)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(actors[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], actors[0])
		return
	}
}

func TestGetActorsSuccess(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()

	repo := &ActorRepo{
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
		ExpectQuery(`SELECT id, name, surname, birthday, gender
		FROM actor `).
		WillReturnRows(rows)

	actors, err := repo.GetActors()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(actors[0], expect[0]) {
		t.Errorf("results not match, want %v, have %v", expect[0], actors[0])
		return
	}
}
