package user

import (
	"context"
	"encoding/json"
	"github.com/fahrurben/gopress/helper/controller"
	"net/http"
)

type Service interface {
	Save(ctx context.Context, request CreateUserRequest) (*int64, error)
}

type Handler struct {
	service Service
}

func CreateHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var createUserRequest CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.Save(r.Context(), createUserRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	controller.WriteSuccessResponse(w, id)
}
