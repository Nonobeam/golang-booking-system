package entity

import (
	"time"
)

type User struct {
	ID        string    `json:"id" db:"id" validate:"required"`
	Name      string    `json:"name" db:"name" validate:"required,min=1,max=255"`
	Email     string    `json:"email" db:"email" validate:"required,email,max=255"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
