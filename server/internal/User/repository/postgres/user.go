package postgres

import (
	"database/sql"
	"server/server/internal/User/repository"
	"server/server/internal/domain/entity"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepositoryI {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) FindUserByUsername(value string) (*entity.User, error) {
	user := &entity.User{}
	row := repo.DB.QueryRow("SELECT id, name, password, status FROM users WHERE name = $1", value)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) FindUserByID(id uint) (*entity.User, error) {
	user := &entity.User{}
	row := repo.DB.QueryRow("SELECT id, name, password, status FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
