package usecase

import (
	"fmt"
	"regexp"
	actorRep "server/server/internal/Actor/repository"
	"server/server/internal/domain/entity"
)

type ActorUsecaseI interface {
	CreateActor(newActor *entity.Actor) (uint, error)
	UpdateActor(newActor *entity.Actor) error
	DeleteActor(id uint) error
}

type ActorUsecase struct {
	actorRepo actorRep.ActorRepositoryI
}

func NewActorUsecase(actorRepI actorRep.ActorRepositoryI) *ActorUsecase {
	return &ActorUsecase{
		actorRepo: actorRepI,
	}
}

func (au ActorUsecase) CreateActor(newActor *entity.Actor) (uint, error) {
	err := au.checkFields(newActor)
	if err != nil {
		return 0, err
	}

	actorId, err := au.actorRepo.CreateActor(newActor)
	if err != nil {
		return 0, err
	}

	return actorId, nil
}

func (au ActorUsecase) UpdateActor(newActor *entity.Actor) error {
	actor, err := au.actorRepo.GetActorById(newActor.ID)
	if err != nil {
		return err
	}

	if actor != nil {
		if newActor.Name != "" {
			actor.Name = newActor.Name
		}

		if newActor.Surname != "" {
			actor.Surname = newActor.Surname
		}

		if newActor.Birthday != "" {
			actor.Birthday = newActor.Birthday
		}

		if newActor.Gender != "" {
			actor.Gender = newActor.Gender
		}

		err = au.checkFields(actor)
		if err != nil {
			return err
		}

		return au.actorRepo.UpdateActor(actor)
	}

	return entity.ErrNotFound
}

func (au ActorUsecase) checkFields(newActor *entity.Actor) error {
	re := regexp.MustCompile(`\d{4}-\d{1,2}-\d{1,2}`)
	if !re.MatchString(newActor.Birthday) {
		fmt.Println(newActor.Birthday)
		return entity.ErrInvalidBirthday
	}

	if !(newActor.Gender == "ж" || newActor.Gender == "м") {
		return entity.ErrInvalidGender
	}

	return nil
}

func (au ActorUsecase) DeleteActor(id uint) error {
	err := au.actorRepo.DeleteActor(id)
	if err != nil {
		return err
	}
	return nil
}
