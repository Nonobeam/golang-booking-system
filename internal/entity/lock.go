package entity

import (
	"time"
)

type Lock struct {
	Key   string        `json:"key" validate:"required"`
	Value string        `json:"value" validate:"required"`
	TTL   time.Duration `json:"ttl" validate:"required"`
}
