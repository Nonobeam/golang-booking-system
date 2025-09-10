package entity

import "time"

type PaymentMessage struct {
	BookingID string    `json:"booking_id" validate:"required"`
	UserID    string    `json:"user_id" validate:"required"`
	EventID   string    `json:"event_id" validate:"required"`
	Amount    float64   `json:"amount" validate:"required,min=0"`
	CreatedAt time.Time `json:"created_at"`
}

type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "PENDING"
	PaymentStatusCompleted PaymentStatus = "COMPLETED"
	PaymentStatusFailed    PaymentStatus = "FAILED"
)
