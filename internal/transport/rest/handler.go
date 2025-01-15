package rest

import (
	"encoding/json"
	"fmt"
	"github.com/dankru/practice-task1/internal/domain"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type UserService interface {
	GetUsers() ([]domain.User, error)
	CreateUser(user domain.User) error
}

type Handler struct {
	service UserService
}

func NewHandler(service UserService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	h.initUserRoutes(r)
	return r
}

func (h *Handler) initUserRoutes(router *mux.Router) {
	users := router.PathPrefix("/users").Subrouter()
	{
		users.HandleFunc("", h.getUsers).Methods(http.MethodGet)
		users.HandleFunc("", h.createUser).Methods(http.MethodPost)
	}
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get users: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(users)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshall users: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(resp))
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var user domain.User
	if err = json.Unmarshal(reqBytes, &user); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err = h.service.CreateUser(user); err != nil {
		http.Error(w, fmt.Sprintf("failed to create user: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
