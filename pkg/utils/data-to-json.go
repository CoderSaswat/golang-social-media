package utils

import (
	"encoding/json"
	"net/http"
	"social-media/pkg/dto"
)

func JsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func SuccessResponse(w http.ResponseWriter, data any, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := dto.NewResponse(message, false, data)
	jsonResponse, _ := json.Marshal(response)
	_, err := w.Write(jsonResponse)
	if err != nil {
		return
	}
}
func ErrorResponse(w http.ResponseWriter, data any, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := dto.NewResponse(err.Error(), true, data)
	jsonResponse, _ := json.Marshal(response)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return
	}
}
