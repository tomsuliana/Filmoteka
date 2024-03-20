package usecase

import (
	mockA "server/internal/Actor/repository/mock_repository"
	mockF "server/internal/Film/repository/mock_repository"
	"server/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewFilmUsecase(mockFilm, mockAct)

	filmwithactors := &entity.FilmWithActors{
		ID:          1,
		Name:        "Terminator 2",
		Description: "Cool film",
		ReleaseDate: "1985-08-22",
		Rating:      5.9,
		Actors: []*entity.Actor{
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
		},
	}

	mockFilm.EXPECT().CreateFilm(entity.ToFilm(filmwithactors)).Return(filmwithactors.ID, nil)
	mockAct.EXPECT().GetActorByName(filmwithactors.Actors[0].Name, filmwithactors.Actors[0].Surname).Return(uint(1), nil)
	mockFilm.EXPECT().AddActorToFilm(filmwithactors.Actors[0].ID, filmwithactors.ID).Return(nil)
	mockAct.EXPECT().GetActorByName(filmwithactors.Actors[1].Name, filmwithactors.Actors[1].Surname).Return(uint(2), nil)
	mockFilm.EXPECT().AddActorToFilm(filmwithactors.Actors[1].ID, filmwithactors.ID).Return(nil)

	actual, err := usecase.CreateFilm(filmwithactors)

	assert.Equal(t, uint(1), actual)
	assert.Nil(t, err)
}

func TestUpdateFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewFilmUsecase(mockFilm, mockAct)

	film := &entity.Film{
		ID:          1,
		Name:        "Terminator 2",
		Description: "Cool film",
		ReleaseDate: "1985-08-22",
		Rating:      5.9,
	}

	mockFilm.EXPECT().GetFilmById(film.ID).Return(film, nil)
	mockFilm.EXPECT().UpdateFilm(film).Return(nil)

	err := usecase.UpdateFilm(film)

	assert.Nil(t, err)
}

func TestAddActorToFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewFilmUsecase(mockFilm, mockAct)

	actor := &entity.Actor{
		ID:       1,
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	filmId := 1

	mockAct.EXPECT().GetActorByName(actor.Name, actor.Surname).Return(actor.ID, nil)
	mockFilm.EXPECT().AddActorToFilm(actor.ID, uint(filmId)).Return(nil)

	err := usecase.AddActorToFilm(actor, uint(filmId))

	assert.Nil(t, err)
}

func TestDeleteActorFromFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewFilmUsecase(mockFilm, mockAct)

	actor := &entity.Actor{
		ID:       1,
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	filmId := 1

	mockAct.EXPECT().GetActorByName(actor.Name, actor.Surname).Return(actor.ID, nil)
	mockFilm.EXPECT().DeleteActorFromFilm(actor.ID, uint(filmId)).Return(nil)

	err := usecase.DeleteActorFromFilm(actor, uint(filmId))

	assert.Nil(t, err)
}

func TestDeleteFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewFilmUsecase(mockFilm, mockAct)

	filmId := 1

	mockFilm.EXPECT().DeleteFilm(uint(filmId)).Return(nil)

	err := usecase.DeleteFilm(uint(filmId))

	assert.Nil(t, err)
}

func TestGetActorsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewFilmUsecase(mockFilm, mockAct)

	films := []*entity.Film{
		{
			ID:          1,
			Name:        "Terminator 2",
			Description: "Cool film",
			ReleaseDate: "1990-10-10",
			Rating:      6.2,
		},
		{
			ID:          2,
			Name:        "Terminator 3",
			Description: "Cool film too",
			ReleaseDate: "1992-10-10",
			Rating:      7.2,
		},
	}

	filmswithactors := []*entity.FilmWithActors{
		{
			ID:          1,
			Name:        "Terminator 2",
			Description: "Cool film",
			ReleaseDate: "1990-10-10",
			Rating:      6.2,
			Actors: []*entity.Actor{
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
			},
		},
		{
			ID:          2,
			Name:        "Terminator 3",
			Description: "Cool film too",
			ReleaseDate: "1992-10-10",
			Rating:      7.2,
			Actors: []*entity.Actor{
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
			},
		},
	}

	mockFilm.EXPECT().GetFilms(false, false).Return(films, nil)
	mockFilm.EXPECT().GetActorsByFilm(films[0].ID).Return(filmswithactors[0].Actors, nil)
	mockFilm.EXPECT().GetActorsByFilm(films[1].ID).Return(filmswithactors[1].Actors, nil)
	actual, err := usecase.GetFilms(false, false)
	assert.Equal(t, filmswithactors, actual)
	assert.Nil(t, err)
}

func TestSearchSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewFilmUsecase(mockFilm, mockAct)

	films := []*entity.Film{
		{
			ID:          1,
			Name:        "Terminator 2",
			Description: "Cool film",
			ReleaseDate: "1990-10-10",
			Rating:      6.2,
		},
		{
			ID:          2,
			Name:        "Terminator 3",
			Description: "Cool film too",
			ReleaseDate: "1992-10-10",
			Rating:      7.2,
		},
	}

	actors := []*entity.Actor{
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

	filmswithactors := []*entity.FilmWithActors{
		{
			ID:          1,
			Name:        "Terminator 2",
			Description: "Cool film",
			ReleaseDate: "1990-10-10",
			Rating:      6.2,
			Actors: []*entity.Actor{
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
			},
		},
		{
			ID:          2,
			Name:        "Terminator 3",
			Description: "Cool film too",
			ReleaseDate: "1992-10-10",
			Rating:      7.2,
			Actors: []*entity.Actor{
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
			},
		},
	}

	word := "Term"

	mockFilm.EXPECT().SearchFilms(word).Return(films, nil)
	mockAct.EXPECT().SearchActors(word).Return(actors, nil)
	mockFilm.EXPECT().GetFilmsByActor(actors[0]).Return(films, nil)
	mockFilm.EXPECT().GetFilmsByActor(actors[1]).Return(films, nil)
	mockFilm.EXPECT().GetActorsByFilm(films[0].ID).Return(actors, nil)
	mockFilm.EXPECT().GetActorsByFilm(films[1].ID).Return(actors, nil)
	actual, err := usecase.Search(word)
	assert.Equal(t, filmswithactors, actual)
	assert.Nil(t, err)
}
