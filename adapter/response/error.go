package response

import (
	"encoding/json"
	"net/http"
)

type exception struct {
	statusCode int
	Errors     []string `json:"errors"`
}

func NewError(statusCode int, err error) *exception {
	return &exception{statusCode: statusCode, Errors: []string{err.Error()}}
}

func (e *exception) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	_ = json.NewEncoder(w).Encode(e)
}
