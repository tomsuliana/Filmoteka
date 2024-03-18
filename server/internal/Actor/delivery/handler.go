package delivery

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	actorUsecase "server/server/internal/Actor/usecase"
	"server/server/internal/domain/entity"
	mw "server/server/internal/middleware"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type ActorHandler struct {
	actors actorUsecase.ActorUsecaseI
	logger *mw.ACLog
}

func NewActorHandler(actors actorUsecase.ActorUsecaseI, logger *mw.ACLog) *ActorHandler {
	return &ActorHandler{
		actors: actors,
		logger: logger,
	}
}

func (handler *ActorHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/actors", handler.CreateActor).Methods(http.MethodPost)
	router.HandleFunc("/api/actors", handler.GetActorList).Methods(http.MethodGet)
	router.HandleFunc("/api/actors/{id:[0-9]+}", handler.UpdateActor).Methods(http.MethodPatch)
	router.HandleFunc("/api/actors/{id:[0-9]+}", handler.DeleteActor).Methods(http.MethodDelete)
}

func (handler *ActorHandler) CreateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqActor := entity.Actor{}

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonbody, &reqActor)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := handler.actors.CreateActor(&reqActor)
	if err != nil {
		handler.logger.LogError("problems with creating actor", err, w.Header().Get("request-id"), r.URL.Path)
		//fmt.Println(err)
		if err == entity.ErrInvalidBirthday {
			w.WriteHeader(http.StatusBadRequest)
		} else if err == entity.ErrInvalidGender {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)

	body := &entity.RespID{ID: id}

	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err != nil {
		handler.logger.LogError("problems marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *ActorHandler) UpdateActor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("id is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("id is not number"), w.Header().Get("request-id"), r.URL.Path)
		return
	}

	id := uint(id64)

	jsonbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateActor := &entity.Actor{ID: id}
	err = json.Unmarshal(jsonbody, &updateActor)
	if err != nil {
		handler.logger.LogError("prbolems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.actors.UpdateActor(updateActor)
	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("actor not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		fmt.Println(err)
		if err == entity.ErrInvalidBirthday {
			w.WriteHeader(http.StatusBadRequest)
		} else if err == entity.ErrInvalidGender {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		handler.logger.LogError("problems updating actor", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *ActorHandler) DeleteActor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strid, ok := vars["id"]
	if !ok {
		handler.logger.LogError("problems with parameters", errors.New("id is missing in parameters"), w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id64, err := strconv.ParseUint(strid, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		handler.logger.LogError("problems with parameters", errors.New("id is not number"), w.Header().Get("request-id"), r.URL.Path)
		return
	}

	id := uint(id64)

	err = handler.actors.DeleteActor(id)
	if err != nil {
		handler.logger.LogError("problems deleting person", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *ActorHandler) GetActorList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	actors, err := handler.actors.GetActors()

	if err != nil {
		handler.logger.LogError("problems with getting actors", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := actors

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
