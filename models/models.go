package models

import (
	"time"
)

type User struct {
	ID int `gorm:"primaryKey;`
	CreatedAt time.Time

	Nume string
	Prenume string
	Email string
	Password string
}