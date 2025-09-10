// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"time"

	"booking-app/internal/entity"
)

type (
	EventService interface {
		Create(ctx context.Context, event *entity.Event) error
		GetByID(ctx context.Context, id string) (*entity.Event, error)
		Update(ctx context.Context, event *entity.Event) error
		Delete(ctx context.Context, id string) error
		List(ctx context.Context) ([]*entity.Event, error)
		GetStatistics(ctx context.Context, eventID string) (*entity.EventStatistics, error)
	}

	UserService interface {
		Create(ctx context.Context, user *entity.User) error
		GetByID(ctx context.Context, id string) (*entity.User, error)
		Update(ctx context.Context, user *entity.User) error
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
	}

	BookingService interface {
		CreateBooking(ctx context.Context, req *entity.CreateBookingRequest) (*entity.Booking, error)
		GetBooking(ctx context.Context, id string) (*entity.Booking, error)
		GetUserBookings(ctx context.Context, userID string) ([]*entity.Booking, error)
		ProcessPayment(ctx context.Context, bookingID string) error
		CancelExpiredBookings(ctx context.Context) error
	}

	LockService interface {
		AcquireBookingLock(ctx context.Context, eventID string, timeout time.Duration) (*entity.Lock, error)
		ReleaseLock(ctx context.Context, lock *entity.Lock) error
	}
)
