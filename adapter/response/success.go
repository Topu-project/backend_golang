package response

import (
	"encoding/json"
	"net/http"
)

type success struct {
	statusCode int
	result     interface{}
}

func NewSuccess(statusCode int, result interface{}) *success {
	return &success{
		statusCode: statusCode,
		result:     result,
	}
}

func (s *success) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s.statusCode)
	json.NewEncoder(w).Encode(s.result)
}
