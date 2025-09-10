package dto

import "time"

type UserRequest struct {
	Name  string
	Email string
}

type UserResponse struct {
	Id           int
	Name         string
	Email        string
	RegisteredAt time.Time `copier:"CreatedAt"`
	UpdatedAt    time.Time
}
type SignUpRequest struct {
	Name     string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}
