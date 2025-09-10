package repo

import (
	"context"
	"database/sql"

	"booking-app/internal/entity"
)

type (
	EventRepo interface {
		Create(ctx context.Context, event *entity.Event) error
		GetByID(ctx context.Context, id string) (*entity.Event, error)
		Update(ctx context.Context, event *entity.Event) error
		Delete(ctx context.Context, id string) error
		List(ctx context.Context) ([]*entity.Event, error)
		GetAvailableTickets(ctx context.Context, eventID string) (int, error)
	}

	UserRepo interface {
		Create(ctx context.Context, user *entity.User) error
		GetByID(ctx context.Context, id string) (*entity.User, error)
		Update(ctx context.Context, user *entity.User) error
		GetByEmail(ctx context.Context, email string) (*entity.User, error)
	}

	BookingRepo interface {
		Create(ctx context.Context, booking *entity.Booking) error
		GetByID(ctx context.Context, id string) (*entity.Booking, error)
		Update(ctx context.Context, booking *entity.Booking) error
		GetByUserID(ctx context.Context, userID string) ([]*entity.Booking, error)
		GetExpiredBookings(ctx context.Context) ([]*entity.Booking, error)
		GetEventStatistics(ctx context.Context, eventID string) (*entity.EventStatistics, error)
		CreateWithLock(ctx context.Context, tx *sql.Tx, booking *entity.Booking, eventID string) error
	}

	LockRepo interface {
		AcquireLock(ctx context.Context, key string, value string, ttl int) (bool, error)
		ReleaseLock(ctx context.Context, key string, value string) error
		ExtendLock(ctx context.Context, key string, value string, ttl int) error
	}
)
