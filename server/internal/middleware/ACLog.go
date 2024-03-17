package middleware

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ACLog struct {
	logger      *zap.SugaredLogger
	errorLogger *zap.SugaredLogger
}

type ResponseRecorder struct {
	http.ResponseWriter
	status int
}

//WriteHeader writes header of response status for AC logs
func (r *ResponseRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

//NewLoggingResponseWriter creates new ResponseRecorder
func NewLoggingResponseWriter(w http.ResponseWriter) *ResponseRecorder {
	return &ResponseRecorder{w, http.StatusOK}
}

//NewACLog createas new object of ACLog
func NewACLog(logger *zap.SugaredLogger, errorLogger *zap.SugaredLogger) *ACLog {
	return &ACLog{
		logger:      logger,
		errorLogger: errorLogger,
	}
}

//ACLogMiddleware creates access logs
func (ac *ACLog) ACLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := NewLoggingResponseWriter(w)
		requestID := uuid.New().String()
		w.Header().Set("request-id", requestID)
		start := time.Now()
		next.ServeHTTP(rec, r)
		status := rec.status
		ac.logger.Infow("Access log info",
			zap.String("method", r.Method),
			zap.String("remote addr", r.RemoteAddr),
			zap.String("url", r.URL.Path),
			zap.String("request-id", requestID),
			zap.Int("status", status),
			zap.Duration("work time", time.Duration(time.Since(start).Microseconds())),
		)
	})
}

//LogError creates error logs
func (ac *ACLog) LogError(message string, err error, requestID string, url string) {
	ac.errorLogger.Errorw(message,
		zap.Error(err),
		zap.String("request-id", requestID),
		zap.String("url", url),
	)
}
