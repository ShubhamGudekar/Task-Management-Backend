package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Status   string `gorm:"not null;default:pending"`
	Priority string `gorm:"not null;default:medium"`
	DueDate  time.Time
	UserID   uint
	User     User
}

const (
	StatusPending  = "pending"
	StatusOngoing  = "ongoing"
	StatusComplete = "complete"

	PriorityLow    = "low"
	PriorityMedium = "medium"
	PriorityHigh   = "high"
)

func IsValidStatus(status string) bool {
	return status == StatusPending || status == StatusOngoing || status == StatusComplete
}

func IsValidPriority(priority string) bool {
	return priority == PriorityLow || priority == PriorityMedium || priority == PriorityHigh
}
