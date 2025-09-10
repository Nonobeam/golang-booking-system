package v1

import (
	"booking-app/pkg/logger"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// NewTranslationRoutes -.
func NewTranslationRoutes(apiV1Group fiber.Router, t interface{}, l logger.Interface) {
	r := &V1{l: l, v: validator.New(validator.WithRequiredStructEnabled())}

	translationGroup := apiV1Group.Group("/translation")

	{
		// TODO: Add actual route handlers when use cases are implemented
		translationGroup.Get("/history", r.placeholderHandler)
		translationGroup.Post("/do-translate", r.placeholderHandler)
	}
}

// placeholderHandler is a temporary handler
func (r *V1) placeholderHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "API endpoint not yet implemented",
		"status":  "placeholder",
	})
}
