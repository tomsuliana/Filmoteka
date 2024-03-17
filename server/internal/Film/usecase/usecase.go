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
	err := fu.checkFields(newFilm)
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

func (fu FilmUsecase) checkFields(newFilm *entity.FilmWithActors) error {
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
