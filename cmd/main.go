package main

import (
	"github.com/dankru/practice-task1/internal/transport/rest"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	handler := rest.NewHandler()
	handler.InitRouter()
	if err := http.ListenAndServe(":8080", handler.InitRouter()); err != nil {
		log.Fatal("Failed to run server", err.Error())
	}
}
