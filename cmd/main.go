package main

import (
	"database/sql"
	"github.com/dankru/practice-task1/internal/repository/psql"
	"github.com/dankru/practice-task1/internal/service"
	"github.com/dankru/practice-task1/internal/transport/rest"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=123")
	if err != nil {
		log.Fatal("DB init failure: ", err.Error())
	}
	defer db.Close()

	userRepo := psql.NewRepository(db)
	userService := service.NewService(userRepo)
	handler := rest.NewHandler(userService)
	handler.InitRouter()

	if err := http.ListenAndServe(":8080", handler.InitRouter()); err != nil {
		log.Fatal("Failed to run server", err.Error())
	}
}
