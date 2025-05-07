package shared

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// JSON menuliskan response JSON standar
func JSON(w http.ResponseWriter, status int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Response{
		Data:    data,
		Message: message,
	})
}

// ErrorResponse menuliskan pesan error
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	JSON(w, status, nil, message)
}
