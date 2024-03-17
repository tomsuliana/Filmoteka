package postgres

import (
	"database/sql"
	"server/server/internal/Actor/repository"
	"server/server/internal/domain/entity"
)

type ActorRepo struct {
	DB *sql.DB
}

func NewActorRepo(db *sql.DB) repository.ActorRepositoryI {
	return &ActorRepo{
		DB: db,
	}
}

func (repo *ActorRepo) CreateActor(actor *entity.Actor) (uint, error) {
	insertActor := `INSERT INTO actor (name, surname, birthday, gender) VALUES ($1, $2, $3, $4) RETURNING ID`
	var Id uint
	err := repo.DB.QueryRow(insertActor, actor.Name, actor.Surname, actor.Birthday, actor.Gender).Scan(&Id)
	if err != nil {
		return 0, err
	}

	return Id, nil
}

func (repo *ActorRepo) UpdateActor(actor *entity.Actor) error {
	updateActor := `UPDATE actor 
				   SET name = $1, surname = $2, birthday = $3, gender = $4
				   WHERE id = $5`
	_, err := repo.DB.Exec(updateActor, actor.Name, actor.Surname, actor.Birthday, actor.Gender, actor.ID)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActorRepo) GetActorById(id uint) (*entity.Actor, error) {
	actor := &entity.Actor{}
	row := repo.DB.QueryRow(`SELECT id, name, surname, birthday, gender FROM actor WHERE id = $1`, id)
	err := row.Scan(&actor.ID, &actor.Name, &actor.Surname, &actor.Birthday, &actor.Gender)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return actor, nil
}

func (repo *ActorRepo) DeleteActor(id uint) error {
	deleteActor := `DELETE FROM actor WHERE id = $1`
	_, err := repo.DB.Exec(deleteActor, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *ActorRepo) GetActorByName(name string, surname string) (uint, error) {
	actor := &entity.Actor{}
	row := repo.DB.QueryRow(`SELECT id FROM actor WHERE name = $1 and surname = $2`, name, surname)
	err := row.Scan(&actor.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	return actor.ID, nil
}
