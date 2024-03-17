package delivery

import (
	"encoding/json"
	//"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	filmUsecase "server/server/internal/Film/usecase"
	"server/server/internal/domain/entity"
	mw "server/server/internal/middleware"
	//"strconv"

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
		fmt.Println(err)
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
