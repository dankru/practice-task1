package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handleUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to run server", err.Error())
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.WriteHeader(http.StatusNotImplemented)
		return
	case http.MethodPost:
		w.WriteHeader(http.StatusNotImplemented)
		return
	case http.MethodPatch:
		w.WriteHeader(http.StatusNotImplemented)
		return
	case http.MethodDelete:
		w.WriteHeader(http.StatusNotImplemented)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}
