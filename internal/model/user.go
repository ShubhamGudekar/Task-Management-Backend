package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
