package v1

import (
	"booking-app/pkg/logger"

	"github.com/go-playground/validator/v10"
)

type V1 struct {
	l logger.Interface
	v *validator.Validate
}
