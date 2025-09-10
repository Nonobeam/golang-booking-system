package entity

import "errors"

type BookingError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

func (e BookingError) Error() string {
	return e.Message
}

const (
	ErrCodeInsufficientTickets = "INSUFFICIENT_TICKETS"
	ErrCodeEventNotFound       = "EVENT_NOT_FOUND"
	ErrCodeUserNotFound        = "USER_NOT_FOUND"
	ErrCodeBookingNotFound     = "BOOKING_NOT_FOUND"
	ErrCodeInvalidInput        = "INVALID_INPUT"
	ErrCodePaymentFailed       = "PAYMENT_FAILED"
	ErrCodeBookingExpired      = "BOOKING_EXPIRED"
	ErrCodeBookingInProgress   = "BOOKING_IN_PROGRESS"
	ErrCodeEventHasBookings    = "EVENT_HAS_BOOKINGS"
)

var (
	ErrInsufficientTickets = BookingError{
		Code:    ErrCodeInsufficientTickets,
		Message: "Insufficient tickets available",
	}
	ErrEventNotFound = BookingError{
		Code:    ErrCodeEventNotFound,
		Message: "Event not found",
	}
	ErrUserNotFound = BookingError{
		Code:    ErrCodeUserNotFound,
		Message: "User not found",
	}
	ErrBookingNotFound = BookingError{
		Code:    ErrCodeBookingNotFound,
		Message: "Booking not found",
	}
	ErrInvalidInput = BookingError{
		Code:    ErrCodeInvalidInput,
		Message: "Invalid input provided",
	}
	ErrPaymentFailed = BookingError{
		Code:    ErrCodePaymentFailed,
		Message: "Payment processing failed",
	}
	ErrBookingExpired = BookingError{
		Code:    ErrCodeBookingExpired,
		Message: "Booking has expired",
	}
	ErrBookingInProgress = BookingError{
		Code:    ErrCodeBookingInProgress,
		Message: "Another booking is in progress for this event",
	}
	ErrEventHasBookings = BookingError{
		Code:    ErrCodeEventHasBookings,
		Message: "Cannot delete event with existing bookings",
	}

	ErrDuplicateEmail = errors.New("email already exists")
	ErrInvalidEmail   = errors.New("invalid email format")
)
