package usecase

import (
	"regexp"
	actorRep "server/internal/Actor/repository"
	filmRep "server/internal/Film/repository"
	"server/internal/domain/entity"
)

type ActorUsecaseI interface {
	CreateActor(newActor *entity.Actor) (uint, error)
	UpdateActor(newActor *entity.Actor) error
	DeleteActor(id uint) error
	GetActors() ([]*entity.ActorWithFilms, error)
}

type ActorUsecase struct {
	actorRepo actorRep.ActorRepositoryI
	filmRepo  filmRep.FilmRepositoryI
}

func NewActorUsecase(actorRepI actorRep.ActorRepositoryI, filmRepI filmRep.FilmRepositoryI) *ActorUsecase {
	return &ActorUsecase{
		actorRepo: actorRepI,
		filmRepo:  filmRepI,
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

func (au ActorUsecase) GetActors() ([]*entity.ActorWithFilms, error) {
	actors, err := au.actorRepo.GetActors()
	if err != nil {
		return nil, err
	}

	actorsWithFilms := []*entity.ActorWithFilms{}
	for _, actor := range actors {
		films, err := au.filmRepo.GetFilmsByActor(actor)
		if err != nil && err != entity.ErrNotFound {
			return nil, err
		}
		actorWithFilms := entity.ToActorWithFilms(actor, films)
		actorsWithFilms = append(actorsWithFilms, actorWithFilms)
	}

	return actorsWithFilms, nil
}
