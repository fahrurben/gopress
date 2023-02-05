package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool
	Error   string
	Data    any
}

func WriteSuccessResponse(w http.ResponseWriter, data any) {
	response := &Response{
		Success: true,
		Data:    data,
	}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func WriteErrorResponse(w http.ResponseWriter, error string) {
	response := &Response{
		Success: false,
		Error:   error,
		Data:    nil,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	w.Write(jsonData)
}
