package dto

import "time"

type UserRequest struct {
	Name  string `binding:"required,min=2,max=100"`
	Email string `binding:"required,email"`
}

type UserResponse struct {
	Id           int
	Name         string
	Email        string
	RegisteredAt time.Time `copier:"CreatedAt"`
	UpdatedAt    time.Time
}
type SignUpRequest struct {
	Name     string `binding:"required,min=2,max=100"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=6,max=32"`
}

type LoginRequest struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=6,max=32"`
}
