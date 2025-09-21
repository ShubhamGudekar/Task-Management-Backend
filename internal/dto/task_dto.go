package dto

import "time"

type TaskRequest struct {
	Title    string    `binding:"required,min=2,max=100"`
	DueDate  time.Time `binding:"required"`
	Priority string    `binding:"oneof low medium high"`
	Status   string    `binding:"oneof pending ongoing complete"`
}

type TaskResponse struct {
	ID        int
	UserId    int
	Title     string
	DueDate   time.Time
	Status    string
	Priority  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
