package usecase

import (
	// "fmt"
	"regexp"
	actorRep "server/server/internal/Actor/repository"
	filmRep "server/server/internal/Film/repository"
	"server/server/internal/domain/entity"
)

type FilmUsecaseI interface {
	CreateFilm(newFilm *entity.FilmWithActors) (uint, error)
	UpdateFilm(newFilm *entity.Film) error
	AddActorToFilm(actor *entity.Actor, filmId uint) error
	DeleteActorFromFilm(actor *entity.Actor, filmId uint) error
	DeleteFilm(filmId uint) error
	GetFilms(name bool, releaseDate bool) ([]*entity.FilmWithActors, error)
}

type FilmUsecase struct {
	filmRepo  filmRep.FilmRepositoryI
	actorRepo actorRep.ActorRepositoryI
}

func NewFilmUsecase(filmRepI filmRep.FilmRepositoryI, actorRepI actorRep.ActorRepositoryI) *FilmUsecase {
	return &FilmUsecase{
		filmRepo:  filmRepI,
		actorRepo: actorRepI,
	}
}

func (fu FilmUsecase) CreateFilm(newFilm *entity.FilmWithActors) (uint, error) {
	err := fu.checkFields(entity.ToFilm(newFilm))
	if err != nil {
		return 0, err
	}

	film := entity.ToFilm(newFilm)

	filmId, err := fu.filmRepo.CreateFilm(film)
	if err != nil {
		return 0, err
	}

	for _, actor := range newFilm.Actors {
		actorId, err := fu.actorRepo.GetActorByName(actor.Name, actor.Surname)
		if err != nil {
			return 0, err
		}

		if actorId != 0 {
			err := fu.filmRepo.AddActorToFilm(actorId, filmId)
			if err != nil {
				return 0, err
			}
		}

	}

	return filmId, nil
}

func (fu FilmUsecase) checkFields(newFilm *entity.Film) error {
	re := regexp.MustCompile(`\d{4}-\d{1,2}-\d{1,2}`)
	if !re.MatchString(newFilm.ReleaseDate) {
		return entity.ErrInvalidReleaseDate
	}

	if newFilm.Rating < 0 || newFilm.Rating > 10 {
		return entity.ErrInvalidRating
	}

	if len(newFilm.Name) > 150 || len(newFilm.Name) < 1 {
		return entity.ErrInvalidName
	}

	if len(newFilm.Description) > 1000 {
		return entity.ErrInvalidDescription
	}

	return nil
}

func (fu FilmUsecase) UpdateFilm(newFilm *entity.Film) error {
	film, err := fu.filmRepo.GetFilmById(newFilm.ID)
	if err != nil {
		return err
	}

	if film != nil {
		if newFilm.Name != "" {
			film.Name = newFilm.Name
		}

		if newFilm.Description != "" {
			film.Description = newFilm.Description
		}

		if newFilm.ReleaseDate != "" {
			film.ReleaseDate = newFilm.ReleaseDate
		}

		if newFilm.Rating != 0 {
			film.Rating = newFilm.Rating
		}

		err = fu.checkFields(film)
		if err != nil {
			return err
		}

		return fu.filmRepo.UpdateFilm(film)
	}

	return entity.ErrNotFound
}

func (fu FilmUsecase) AddActorToFilm(actor *entity.Actor, filmId uint) error {
	actorId, err := fu.actorRepo.GetActorByName(actor.Name, actor.Surname)
	if err != nil {
		return err
	}

	if actorId != 0 {
		err := fu.filmRepo.AddActorToFilm(actorId, filmId)
		if err != nil {
			return err
		}
	} else {
		return entity.ErrNotFound
	}

	return nil
}

func (fu FilmUsecase) DeleteActorFromFilm(actor *entity.Actor, filmId uint) error {
	actorId, err := fu.actorRepo.GetActorByName(actor.Name, actor.Surname)
	if err != nil {
		return err
	}

	if actorId != 0 {
		err := fu.filmRepo.DeleteActorFromFilm(actorId, filmId)
		if err != nil {
			return err
		}
	} else {
		return entity.ErrNotFound
	}

	return nil
}

func (fu FilmUsecase) DeleteFilm(filmId uint) error {
	err := fu.filmRepo.DeleteFilm(filmId)
	if err != nil {
		return err
	}
	return nil
}

func (fu FilmUsecase) GetFilms(name bool, releaseDate bool) ([]*entity.FilmWithActors, error) {
	films, err := fu.filmRepo.GetFilms(name, releaseDate)
	if err != nil {
		return nil, err
	}
	filmsWithActors := []*entity.FilmWithActors{}
	for _, film := range films {
		actors, err := fu.filmRepo.GetActorsByFilm(film.ID)
		if err != nil {
			return nil, err
		}
		filmwithacts := entity.ToFilmWithActors(film, actors)
		filmsWithActors = append(filmsWithActors, filmwithacts)
	}

	return filmsWithActors, nil
}
