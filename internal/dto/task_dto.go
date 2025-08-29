package dto

import "time"

type TaskRequest struct {
	UserId   int
	Title    string
	DueDate  time.Time
	Priority string
	Status   string
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
