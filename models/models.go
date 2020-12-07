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

// Kard is the model that deals with the transfer of cards
type Kard struct {
	ID int `gorm:"primariKey"`
	CreatedAt time.Time

	KardID string `gorm:"index"`
	From string `gorm:"index"`
	To string `gorm:"index"`

	Background string
	Message string
}