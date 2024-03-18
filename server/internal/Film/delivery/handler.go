package delivery

import (
	"encoding/json"
	"errors"
	//"fmt"
	"io/ioutil"
	"net/http"
	filmUsecase "server/server/internal/Film/usecase"
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

type FilmHandler struct {
	films  filmUsecase.FilmUsecaseI
	logger *mw.ACLog
}

func NewFilmHandler(films filmUsecase.FilmUsecaseI, logger *mw.ACLog) *FilmHandler {
	return &FilmHandler{
		films:  films,
		logger: logger,
	}
}

func (handler *FilmHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/films", handler.CreateFilm).Methods(http.MethodPost)
	router.HandleFunc("/api/films", handler.GetFilmList).Methods(http.MethodGet)
	router.HandleFunc("/api/films/{id:[0-9]+}/addactor", handler.AddActorToFilm).Methods(http.MethodPatch)
	router.HandleFunc("/api/films/{id:[0-9]+}/deleteactor", handler.DeleteActorFromFilm).Methods(http.MethodPatch)
	router.HandleFunc("/api/films/{id:[0-9]+}", handler.UpdateFilm).Methods(http.MethodPatch)
	router.HandleFunc("/api/films/{id:[0-9]+}", handler.DeleteFilm).Methods(http.MethodDelete)
	router.HandleFunc("/api/films/", handler.Search).Methods(http.MethodGet)
}

func (handler *FilmHandler) CreateFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqFilm := entity.FilmWithActors{}

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(jsonbody, &reqFilm)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := handler.films.CreateFilm(&reqFilm)
	if err != nil {
		handler.logger.LogError("problems with creating film", err, w.Header().Get("request-id"), r.URL.Path)
		//fmt.Println(err)
		if err == entity.ErrInvalidName || err == entity.ErrInvalidDescription || err == entity.ErrInvalidReleaseDate || err == entity.ErrInvalidRating {
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

func (handler *FilmHandler) UpdateFilm(w http.ResponseWriter, r *http.Request) {
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

	updateFilm := &entity.Film{ID: id}
	err = json.Unmarshal(jsonbody, &updateFilm)
	if err != nil {
		handler.logger.LogError("prbolems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.films.UpdateFilm(updateFilm)
	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("film not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		//fmt.Println(err)
		if err == entity.ErrInvalidName || err == entity.ErrInvalidDescription || err == entity.ErrInvalidReleaseDate || err == entity.ErrInvalidRating {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		handler.logger.LogError("problems updating film", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *FilmHandler) AddActorToFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	err = handler.films.AddActorToFilm(&reqActor, id)
	if err != nil {
		handler.logger.LogError("problems with adding actor", err, w.Header().Get("request-id"), r.URL.Path)
		if err == entity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}

func (handler *FilmHandler) DeleteActorFromFilm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	err = handler.films.DeleteActorFromFilm(&reqActor, id)
	if err != nil {
		handler.logger.LogError("problems with adding actor", err, w.Header().Get("request-id"), r.URL.Path)
		if err == entity.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
}

func (handler *FilmHandler) DeleteFilm(w http.ResponseWriter, r *http.Request) {
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

	err = handler.films.DeleteFilm(id)
	if err != nil {
		if err == entity.ErrNotFound {
			handler.logger.LogError("film not found", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err == entity.ErrInvalidName || err == entity.ErrInvalidDescription || err == entity.ErrInvalidReleaseDate || err == entity.ErrInvalidRating {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

		handler.logger.LogError("problems deleting film", err, w.Header().Get("request-id"), r.URL.Path)
		return
	}
}

func (handler *FilmHandler) GetFilmList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nameParameter := false
	releaseDateParameter := false

	name := r.URL.Query().Get("name")
	releaseDate := r.URL.Query().Get("release_date")

	if name == "true" {
		nameParameter = true
	}

	if releaseDate == "true" {
		releaseDateParameter = true
	}

	films, err := handler.films.GetFilms(nameParameter, releaseDateParameter)

	if err != nil {
		handler.logger.LogError("problems with getting films", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := films

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *FilmHandler) Search(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	search := r.URL.Query().Get("search")

	films, err := handler.films.Search(search)

	if err != nil {
		handler.logger.LogError("problems with getting films", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := films

	encoder := json.NewEncoder(w)
	err = encoder.Encode(&Result{Body: body})

	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
