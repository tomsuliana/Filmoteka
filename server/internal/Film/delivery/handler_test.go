package delivery

import (
	"bytes"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http/httptest"

	"server/config"
	mockF "server/internal/Film/usecase/mock_usecase"
	mw "server/internal/middleware"
	"testing"

	"server/internal/domain/entity"

	"github.com/golang/mock/gomock"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestCreateFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/films"
	mockAct := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockAct, logger)

	film := &entity.FilmWithActors{
		Name:        "Terminator 2",
		Description: "Cool film",
		ReleaseDate: "1990-10-10",
		Rating:      6.2,
	}

	var jsonfilm = map[string]interface{}{
		"Name":        "Terminator 2",
		"Description": "Cool film",
		"ReleaseDate": "1990-10-10",
		"Rating":      6.2,
	}

	var ActorID uint
	ActorID = 1

	mockAct.EXPECT().CreateFilm(film).Return(ActorID, nil)

	body, err := json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateFilm(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestCreateFilmFail(t *testing.T) {
	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/films"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	film := &entity.FilmWithActors{
		Name:        "Terminator 2",
		Description: "Cool film",
		ReleaseDate: "1990-10-10",
		Rating:      6.2,
	}

	var jsonfilm = map[string]interface{}{
		"Name":        "Terminator 2",
		"Description": "Cool film",
		"ReleaseDate": "1990-10-10",
		"Rating":      6.2,
	}

	body, err := json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))

	w := httptest.NewRecorder()

	handler.CreateFilm(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	req = httptest.NewRequest("POST", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.CreateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockFilm.EXPECT().CreateFilm(film).Return(uint(0), entity.ErrInternalServerError)

	body, err = json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.CreateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockFilm.EXPECT().CreateFilm(film).Return(uint(0), entity.ErrInvalidReleaseDate)

	body, err = json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.CreateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockFilm.EXPECT().CreateFilm(film).Return(uint(0), entity.ErrInvalidRating)

	body, err = json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.CreateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestUpdateActorSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/films/1"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	film := &entity.Film{
		ID:   1,
		Name: "Termin 2",
	}

	var jsonfilm = map[string]interface{}{
		"Name": "Termin 2",
	}

	mockFilm.EXPECT().UpdateFilm(film).Return(nil)

	body, err := json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.UpdateFilm(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
}

func TestUpdateFilmFail(t *testing.T) {
	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/films/1"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	film := &entity.Film{
		ID:   1,
		Name: "Termin 2",
	}

	var jsonfilm = map[string]interface{}{
		"Name": "Termin 2",
	}

	body, err := json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))

	w := httptest.NewRecorder()

	handler.UpdateFilm(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

	req = httptest.NewRequest("POST", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.UpdateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

	req = httptest.NewRequest("POST", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars := map[string]string{
		"id": "yy",
	}

	req = mux.SetURLVars(req, vars)

	handler.UpdateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

	mockFilm.EXPECT().UpdateFilm(film).Return(entity.ErrInternalServerError)

	body, err = json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.UpdateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)

	mockFilm.EXPECT().UpdateFilm(film).Return(entity.ErrNotFound)

	body, err = json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.UpdateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 404, resp.StatusCode)

	mockFilm.EXPECT().UpdateFilm(film).Return(entity.ErrInvalidReleaseDate)

	body, err = json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.UpdateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

	mockFilm.EXPECT().UpdateFilm(film).Return(entity.ErrInvalidRating)

	body, err = json.Marshal(jsonfilm)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.UpdateFilm(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

}

func TestAddActorToFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/films/1"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	actor := &entity.Actor{
		Name:    "john",
		Surname: "doe",
	}

	var jsonactor = map[string]interface{}{
		"Name":    "john",
		"Surname": "doe",
	}

	var filmId uint
	filmId = 1

	mockFilm.EXPECT().AddActorToFilm(actor, filmId).Return(nil)

	body, err := json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.AddActorToFilm(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
}

func TestAddActorToFilmFail(t *testing.T) {
	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/films/1"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	actor := &entity.Actor{
		Name:    "john",
		Surname: "doe",
	}

	var jsonactor = map[string]interface{}{
		"Name":    "john",
		"Surname": "doe",
	}

	var filmId uint
	filmId = 1

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(nil))
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.AddActorToFilm(w, req)

	resp := w.Result()

	require.Equal(t, 400, resp.StatusCode)

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.AddActorToFilm(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "7u",
	}

	req = mux.SetURLVars(req, vars)

	handler.AddActorToFilm(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)

	mockFilm.EXPECT().AddActorToFilm(actor, filmId).Return(entity.ErrInternalServerError)

	body, err := json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.AddActorToFilm(w, req)

	resp = w.Result()

	require.Equal(t, 500, resp.StatusCode)

	mockFilm.EXPECT().AddActorToFilm(actor, filmId).Return(entity.ErrNotFound)

	body, err = json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.AddActorToFilm(w, req)

	resp = w.Result()

	require.Equal(t, 404, resp.StatusCode)
}

func TestDeleteActorFromFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/films/1"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	actor := &entity.Actor{
		Name:    "john",
		Surname: "doe",
	}

	var jsonactor = map[string]interface{}{
		"Name":    "john",
		"Surname": "doe",
	}

	var filmId uint
	filmId = 1

	mockFilm.EXPECT().DeleteActorFromFilm(actor, filmId).Return(nil)

	body, err := json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.DeleteActorFromFilm(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
}

func TestDeleteActorFromFilmFail(t *testing.T) {
	baseLogger, err := config.Cfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer baseLogger.Sync()

	errorLogger, err := config.ErrorCfg.Build()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer errorLogger.Sync()
	logger := mw.NewACLog(baseLogger.Sugar(), errorLogger.Sugar())
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	apiPath := "/api/films/1"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	actor := &entity.Actor{
		Name:    "john",
		Surname: "doe",
	}

	var jsonactor = map[string]interface{}{
		"Name":    "john",
		"Surname": "doe",
	}

	var filmId uint
	filmId = 1

	req := httptest.NewRequest("PATCH", apiPath, bytes.NewReader(nil))
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.DeleteActorFromFilm(w, req)

	resp := w.Result()

	require.Equal(t, 400, resp.StatusCode)

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.DeleteActorFromFilm(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "7u",
	}

	req = mux.SetURLVars(req, vars)

	handler.DeleteActorFromFilm(w, req)

	resp = w.Result()

	require.Equal(t, 400, resp.StatusCode)

	mockFilm.EXPECT().DeleteActorFromFilm(actor, filmId).Return(entity.ErrInternalServerError)

	body, err := json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.DeleteActorFromFilm(w, req)

	resp = w.Result()

	require.Equal(t, 500, resp.StatusCode)

	mockFilm.EXPECT().DeleteActorFromFilm(actor, filmId).Return(entity.ErrNotFound)

	body, err = json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req = httptest.NewRequest("PATCH", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	vars = map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.DeleteActorFromFilm(w, req)

	resp = w.Result()

	require.Equal(t, 404, resp.StatusCode)
}

func TestDeleteFilmSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/films/1"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

	var filmId uint
	filmId = 1

	mockFilm.EXPECT().DeleteFilm(filmId).Return(nil)

	req := httptest.NewRequest("DELETE", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.DeleteFilm(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestGetFilmListSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/films"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

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

	mockFilm.EXPECT().GetFilms(false, false).Return(filmswithactors, nil)

	req := httptest.NewRequest("GET", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.GetFilmList(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestSearchSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/films"
	mockFilm := mockF.NewMockFilmUsecaseI(ctrl)
	handler := NewFilmHandler(mockFilm, logger)

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

	word := ""

	mockFilm.EXPECT().Search(word).Return(filmswithactors, nil)

	req := httptest.NewRequest("GET", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	req.URL.Query().Add("search", "")
	w := httptest.NewRecorder()

	handler.Search(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}
