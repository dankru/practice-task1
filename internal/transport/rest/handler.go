package rest

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	return r
}

func NewHandler() *Handler {
	return &Handler{}
}
