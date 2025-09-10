package entity

import (
	"time"
)

type BookingStatus string

const (
	BookingStatusPending   BookingStatus = "PENDING"
	BookingStatusConfirmed BookingStatus = "CONFIRMED"
	BookingStatusCancelled BookingStatus = "CANCELLED"
)

type Booking struct {
	ID          string        `json:"id" db:"id" validate:"required"`
	UserID      string        `json:"user_id" db:"user_id" validate:"required"`
	EventID     string        `json:"event_id" db:"event_id" validate:"required"`
	TicketCount int           `json:"ticket_count" db:"ticket_count" validate:"required,min=1"`
	TotalPrice  float64       `json:"total_price" db:"total_price" validate:"required,min=0"`
	Status      BookingStatus `json:"status" db:"status" validate:"required,oneof=PENDING CONFIRMED CANCELLED"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
	ExpiresAt   time.Time     `json:"expires_at" db:"expires_at" validate:"required"`
}

type CreateBookingRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	EventID     string `json:"event_id" validate:"required"`
	TicketCount int    `json:"ticket_count" validate:"required,min=1"`
}
