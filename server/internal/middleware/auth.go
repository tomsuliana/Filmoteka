package middleware

import (
	"fmt"
	"net/http"
	sessionUsecase "server/internal/Session/usecase"
	"time"
)

type SessionMiddleware struct {
	sessionUC sessionUsecase.SessionUsecaseI
	logger    *ACLog
}

func NewSessionMiddleware(sessionUC sessionUsecase.SessionUsecaseI, logger *ACLog) *SessionMiddleware {
	return &SessionMiddleware{
		sessionUC: sessionUC,
		logger:    logger,
	}
}

func (mw *SessionMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		cookie, err := r.Cookie("session_id")

		if err == http.ErrNoCookie {
			mw.logger.LogError("no cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			return
		} else if err != nil {
			mw.logger.LogError("problems with getting cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
		}

		userId, err := mw.sessionUC.Check(cookie.Value)
		fmt.Println(userId)
		if err != nil {
			mw.logger.LogError("problems with getting user by cookie", err, w.Header().Get("request-id"), r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if userId == 0 {
			mw.logger.LogError("user not found", err, w.Header().Get("request-id"), r.URL.Path)
			cookie.Expires = time.Now().AddDate(0, 0, -1)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if r.Method != http.MethodGet && r.URL.Path != "/api/logout" {
			user, err := mw.sessionUC.GetUserProfile(cookie.Value)
			if err != nil {
				mw.logger.LogError("problems with getting user by cookie", err, w.Header().Get("request-id"), r.URL.Path)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if user.Status != "admin" {
				mw.logger.LogError("not admin", err, w.Header().Get("request-id"), r.URL.Path)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
