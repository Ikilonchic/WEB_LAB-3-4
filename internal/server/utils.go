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

// Error ...
func Error(res http.ResponseWriter, code int, err error) {
	Respond(res, code, map[string]interface{}{"error": err.Error()})
}

// Message ...
func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

// Respond ...
func Respond(res http.ResponseWriter, code int, data map[string]interface{})  {
	res.WriteHeader(code)
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	res.Header().Add("Content-Type", "application/json")
	if data != nil {
		json.NewEncoder(res).Encode(data)
	}
}