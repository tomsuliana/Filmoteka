package repository

import (
	"server/server/internal/domain/entity"
)

//SessionRepositoryI interface
type SessionRepositoryI interface {
	Create(cookie *entity.Cookie) error
	Check(sessionToken string) (*entity.Cookie, error)
	Delete(cookie *entity.DBDeleteCookie) error
}
