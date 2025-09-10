package entity

import (
	"time"
)

type Event struct {
	ID           string    `json:"id" db:"id" validate:"required"`
	Name         string    `json:"name" db:"name" validate:"required,min=1,max=255"`
	Description  string    `json:"description" db:"description" validate:"max=1000"`
	DateTime     time.Time `json:"date_time" db:"date_time" validate:"required"`
	TotalTickets int       `json:"total_tickets" db:"total_tickets" validate:"required,min=1"`
	TicketPrice  float64   `json:"ticket_price" db:"ticket_price" validate:"required,min=0"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type EventStatistics struct {
	EventID          string  `json:"event_id"`
	TotalTickets     int     `json:"total_tickets"`
	SoldTickets      int     `json:"sold_tickets"`
	Revenue          float64 `json:"revenue"`
	AvailableTickets int     `json:"available_tickets"`
}
