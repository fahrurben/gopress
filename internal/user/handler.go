package user

import (
	"encoding/json"
	"github.com/fahrurben/gopress/helper/controller"
	"net/http"
)

type Handler struct {
	userRepository Repository
}

func CreateHandler(userRepository Repository) *Handler {
	return &Handler{userRepository: userRepository}
}

func (h *Handler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.userRepository.Save(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	controller.WriteSuccessResponse(w, id)
}
