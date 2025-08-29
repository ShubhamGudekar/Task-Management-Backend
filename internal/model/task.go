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

func (t Task) IsValidStatus() bool {
	return t.Status == StatusPending || t.Status == StatusOngoing || t.Status == StatusComplete
}

func (t Task) IsValidPriority() bool {
	return t.Priority == PriorityLow || t.Priority == PriorityMedium || t.Priority == PriorityHigh
}
