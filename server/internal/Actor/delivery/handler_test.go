package delivery

import (
	"bytes"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http/httptest"

	"server/config"
	mockA "server/internal/Actor/usecase/mock_usecase"
	mw "server/internal/middleware"
	"testing"

	"server/internal/domain/entity"

	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/require"

	"github.com/gorilla/mux"
)

func TestCreateActorSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/actors"
	mockAct := mockA.NewMockActorUsecaseI(ctrl)
	handler := NewActorHandler(mockAct, logger)

	actor := &entity.Actor{
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	var jsonactor = map[string]interface{}{
		"Name":     "john",
		"Surname":  "doe",
		"Birthday": "1985-08-22",
		"Gender":   "ж",
	}

	var ActorID uint
	ActorID = 1

	mockAct.EXPECT().CreateActor(actor).Return(ActorID, nil)

	body, err := json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	handler.CreateActor(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 201, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))
}

func TestCreateActorFail(t *testing.T) {
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
	apiPath := "/api/actors"
	mockAct := mockA.NewMockActorUsecaseI(ctrl)
	handler := NewActorHandler(mockAct, logger)

	actor := &entity.Actor{
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	var jsonactor = map[string]interface{}{
		"Name":     "john",
		"Surname":  "doe",
		"Birthday": "1985-08-22",
		"Gender":   "ж",
	}

	body, err := json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))

	w := httptest.NewRecorder()

	handler.CreateActor(w, req)

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

	handler.CreateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockAct.EXPECT().CreateActor(actor).Return(uint(0), entity.ErrInternalServerError)

	body, err = json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.CreateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockAct.EXPECT().CreateActor(actor).Return(uint(0), entity.ErrInvalidBirthday)

	body, err = json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.CreateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

	mockAct.EXPECT().CreateActor(actor).Return(uint(0), entity.ErrInvalidGender)

	body, err = json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req = httptest.NewRequest("POST", apiPath, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.CreateActor(w, req)

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
	apiPath := "/api/actors/1"
	mockAct := mockA.NewMockActorUsecaseI(ctrl)
	handler := NewActorHandler(mockAct, logger)

	actor := &entity.Actor{
		ID:   1,
		Name: "johnn",
	}

	var jsonactor = map[string]interface{}{
		"Name": "johnn",
	}

	mockAct.EXPECT().UpdateActor(actor).Return(nil)

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

	handler.UpdateActor(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 200, resp.StatusCode)
}

func TestUpdateActorFail(t *testing.T) {
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
	apiPath := "/api/actors/1"
	mockAct := mockA.NewMockActorUsecaseI(ctrl)
	handler := NewActorHandler(mockAct, logger)

	actor := &entity.Actor{
		ID:       1,
		Name:     "john",
		Surname:  "doe",
		Birthday: "1985-08-22",
		Gender:   "ж",
	}

	var jsonactor = map[string]interface{}{
		"Name":     "john",
		"Surname":  "doe",
		"Birthday": "1985-08-22",
		"Gender":   "ж",
	}

	body, err := json.Marshal(jsonactor)
	if err != nil {
		return
	}

	req := httptest.NewRequest("POST", apiPath, bytes.NewReader(body))

	w := httptest.NewRecorder()

	handler.UpdateActor(w, req)

	resp := w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

	req = httptest.NewRequest("POST", apiPath, nil)
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	handler.UpdateActor(w, req)

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

	handler.UpdateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

	mockAct.EXPECT().UpdateActor(actor).Return(entity.ErrInternalServerError)

	body, err = json.Marshal(jsonactor)
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

	handler.UpdateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 500, resp.StatusCode)

	mockAct.EXPECT().UpdateActor(actor).Return(entity.ErrNotFound)

	body, err = json.Marshal(jsonactor)
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

	handler.UpdateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 404, resp.StatusCode)

	mockAct.EXPECT().UpdateActor(actor).Return(entity.ErrInvalidBirthday)

	body, err = json.Marshal(jsonactor)
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

	handler.UpdateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

	mockAct.EXPECT().UpdateActor(actor).Return(entity.ErrInvalidGender)

	body, err = json.Marshal(jsonactor)
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

	handler.UpdateActor(w, req)

	resp = w.Result()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	require.Equal(t, 400, resp.StatusCode)

}

func TestDeleteActorSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/actors/1"
	mockAct := mockA.NewMockActorUsecaseI(ctrl)
	handler := NewActorHandler(mockAct, logger)

	var ActorID uint
	ActorID = 1

	mockAct.EXPECT().DeleteActor(ActorID).Return(nil)

	req := httptest.NewRequest("DELETE", apiPath, bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "1",
	}

	req = mux.SetURLVars(req, vars)

	handler.DeleteActor(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
}

func TestGetActorsListSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var logger *mw.ACLog
	apiPath := "/api/actors"
	mock := mockA.NewMockActorUsecaseI(ctrl)
	handler := NewActorHandler(mock, logger)

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

	mock.EXPECT().GetActors().Return(actorswithfilms, nil)

	req := httptest.NewRequest("GET", apiPath, nil)
	w := httptest.NewRecorder()

	handler.GetActorList(w, req)

	resp := w.Result()

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, "application/json", resp.Header.Get("Content-Type"))

}
