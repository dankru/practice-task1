package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string
	RegisteredAt time.Time
}

func main() {

	http.HandleFunc("/users", handleUsers)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed to run server", err.Error())
	}
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=123")
	if err != nil {
		log.Fatal("DB initialization failure: ", err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			log.Fatal(err.Error())
		}
		users = append(users, u)
	}

	resp, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err.Error())
	}
	w.Write(resp)
}
