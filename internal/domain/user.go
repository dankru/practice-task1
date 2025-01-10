package domain

import "time"

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string
	RegisteredAt time.Time
}
