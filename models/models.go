package models

import (
	"time"
)

// User is the model that deals with authentication
type User struct {
	ID int `gorm:"primaryKey"`
	CreatedAt time.Time

	Nume string
	Prenume string
	Email string
	Password string
}