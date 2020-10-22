package server

import (
	"encoding/json"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	code int
}

func (res *responseWriter) WriteHeader(statusCode int) {
	res.code = statusCode
	res.ResponseWriter.WriteHeader(statusCode)
}

func (serve *Server) error(res http.ResponseWriter, req *http.Request, code int, err error) {
	serve.respond(res, req, code, map[string]string{"error": err.Error()})
}

func (serve *Server) respond(res http.ResponseWriter, req *http.Request, code int, data interface{}) {
	res.WriteHeader(code)
	if data != nil {
		json.NewEncoder(res).Encode(data)
	}
}