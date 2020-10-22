package server

import (
	"time"
	"net/http"
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

func (serve *Server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		id := uuid.New().String()
		res.Header().Set("X-Request-ID", id)
		next.ServeHTTP(res, req.WithContext(context.WithValue(req.Context(), ctxKeyRequestID, id)))
	})
}

func (serve *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := serve.logger.WithFields(logrus.Fields{
			"remote_addr": req.RemoteAddr,
			"request_id":  req.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", req.Method, req.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, req)

		var level logrus.Level
		
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}

		logger.Logf(level, "completed with %d %s in %v", rw.code, http.StatusText(rw.code), time.Now().Sub(start))
	})
}