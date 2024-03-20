package usecase

import (
	mockA "server/internal/Actor/repository/mock_repository"
	mockF "server/internal/Film/repository/mock_repository"
	"server/internal/domain/entity"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateActorSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewActorUsecase(mockAct, mockFilm)

	actor := &entity.Actor{
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	mockAct.EXPECT().CreateActor(actor).Return(uint(1), nil)
	actual, err := usecase.CreateActor(actor)
	assert.Equal(t, uint(1), actual)
	assert.Nil(t, err)
}

func TestUpdateActorSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewActorUsecase(mockAct, mockFilm)

	actor := &entity.Actor{
		ID:       1,
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	mockAct.EXPECT().GetActorById(actor.ID).Return(actor, nil)
	mockAct.EXPECT().UpdateActor(actor).Return(nil)
	err := usecase.UpdateActor(actor)
	assert.Nil(t, err)
}

func TestDeleteActorSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewActorUsecase(mockAct, mockFilm)

	var id uint
	id = 1

	mockAct.EXPECT().DeleteActor(id).Return(nil)
	err := usecase.DeleteActor(id)
	assert.Nil(t, err)
}

func TestGetActorsSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAct := mockA.NewMockActorRepositoryI(ctrl)
	mockFilm := mockF.NewMockFilmRepositoryI(ctrl)
	usecase := NewActorUsecase(mockAct, mockFilm)

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

	actorswithfilms := []*entity.ActorWithFilms{
		{
			ID:       1,
			Name:     "john",
			Surname:  "doe",
			Birthday: "1985-08-22",
			Gender:   "ж",
			Films: []*entity.Film{
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
			},
		},
		{
			ID:       2,
			Name:     "john",
			Surname:  "doe",
			Birthday: "1985-08-22",
			Gender:   "ж",
			Films: []*entity.Film{
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
			},
		},
	}

	mockAct.EXPECT().GetActors().Return(actors, nil)
	mockFilm.EXPECT().GetFilmsByActor(actors[0]).Return(actorswithfilms[0].Films, nil)
	mockFilm.EXPECT().GetFilmsByActor(actors[1]).Return(actorswithfilms[1].Films, nil)
	actual, err := usecase.GetActors()
	assert.Equal(t, actorswithfilms, actual)
	assert.Nil(t, err)
}
