package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	sessionUsecase "server/internal/Session/usecase"
	"server/internal/domain/entity"
	mw "server/internal/middleware"
	"time"
)

type Result struct {
	Body interface{}
}

type RespError struct {
	Err string
}

type SessionHandler struct {
	sessions sessionUsecase.SessionUsecaseI
	logger   *mw.ACLog
}

func NewSessionHandler(sessions sessionUsecase.SessionUsecaseI, logger *mw.ACLog) *SessionHandler {
	return &SessionHandler{
		sessions: sessions,
		logger:   logger,
	}
}

func (handler *SessionHandler) RegisterAuthHandler(router *mux.Router) {
	router.HandleFunc("/api/logout", handler.Logout).Methods(http.MethodDelete)
}

func (handler *SessionHandler) RegisterHandler(router *mux.Router) {
	router.HandleFunc("/api/login", handler.Login).Methods(http.MethodPost)
}

// Login godoc
// @Summary      Log in user
// @Description  Log in user
// @Tags        users
// @Accept     application/json
// @Produce  application/json
// @Param    user body User true "user object for login"
// @Success  200 {object}  string "success login User return cookie"
// @Failure 400 {object} error "bad request"
// @Failure 404 {object} error "not found"
// @Failure 500 {object} error "internal server error"
// @Router   /api/login [post]
func (handler *SessionHandler) Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Content-Type", "application/json")
	if r.Header.Get("Content-Type") != "application/json" {
		handler.logger.LogError("bad content-type", entity.ErrBadContentType, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonbody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		handler.logger.LogError("problems with reading json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reqUser := entity.User{}
	err = json.Unmarshal(jsonbody, &reqUser)

	if err != nil {
		handler.logger.LogError("problems with unmarshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cookieUC, err := handler.sessions.Login(&reqUser)

	if err != nil {
		if err == entity.ErrInternalServerError {
			handler.logger.LogError("problems with creating cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			handler.logger.LogError("incorrect data", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	}

	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    cookieUC.SessionToken,
		Expires:  time.Now().Add(cookieUC.MaxAge),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, cookie)

	user, err := handler.sessions.GetUserProfile(cookie.Value)
	if err == entity.ErrInternalServerError {
		handler.logger.LogError("problems with getting profile", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&Result{Body: user})
	if err != nil {
		handler.logger.LogError("problems with marshalling json", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Logout godoc
// @Summary      Log out user
// @Description  Log out user
// @Tags        users
// @Param    cookie header string true "Log out user"
// @Success 200 "void" "success log out"
// @Failure 400 {object} error "bad request"
// @Failure 401 {object} error "unauthorized"
// @Router   /api/logout [delete]
func (handler *SessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("session_id")
	err := handler.sessions.Logout(&entity.Cookie{
		SessionToken: cookie.Value,
	})

	if err != nil {
		handler.logger.LogError("problems with deleting cookie", err, w.Header().Get("request-id"), r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie.Expires = time.Now().AddDate(0, 0, -1)
	http.SetCookie(w, cookie)
}
