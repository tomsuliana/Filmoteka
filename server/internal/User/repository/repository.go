package repository

import (
	"server/internal/domain/entity"
)

type UserRepositoryI interface {
	FindUserByID(id uint) (*entity.User, error)
	FindUserByUsername(value string) (*entity.User, error)
}
