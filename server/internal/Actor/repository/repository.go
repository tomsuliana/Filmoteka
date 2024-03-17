package repository

import (
	"server/server/internal/domain/entity"
)

type ActorRepositoryI interface {
	CreateActor(actor *entity.Actor) (uint, error)
	UpdateActor(actor *entity.Actor) error
	GetActorById(id uint) (*entity.Actor, error)
	DeleteActor(id uint) error
	GetActorByName(name string, surname string) (uint, error)
}
