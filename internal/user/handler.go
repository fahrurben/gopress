package user

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/fahrurben/gopress/helper/controller"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type Service interface {
	Save(ctx context.Context, request CreateUserRequest) (*int64, error)
	Update(ctx context.Context, id int, request UpdateUserRequest) (bool, error)
	Delete(ctx context.Context, id int) error
	FindAll(int, int) ([]User, int, int, error)
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
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	id, err := h.service.Save(r.Context(), createUserRequest)

	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	controller.WriteSuccessResponse(w, id)
}

func (h *Handler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	var updateUserRequest UpdateUserRequest
	err = json.NewDecoder(r.Body).Decode(&updateUserRequest)
	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	_, err = h.service.Update(r.Context(), id, updateUserRequest)

	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	controller.WriteSuccessResponse(w, id)
}

func (h *Handler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	err = h.service.Delete(r.Context(), id)

	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	controller.WriteSuccessResponse(w, id)
}

func (h *Handler) SelectUserHandler(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(chi.URLParam(r, "page"))
	limit, err := strconv.Atoi(chi.URLParam(r, "limit"))
	fmt.Println(page, limit)

	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	users, totalCount, totalPage, err := h.service.FindAll(page, limit)

	if err != nil {
		controller.WriteErrorResponse(w, err.Error())
		return
	}

	w.Header().Set("Pagination-Total-Count", strconv.Itoa(totalCount))
	w.Header().Set("Pagination-Total-Page", strconv.Itoa(totalPage))
	controller.WriteSuccessResponse(w, users)
}
