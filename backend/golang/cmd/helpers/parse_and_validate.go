package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var v validator.Validate = *validator.New()

func ParseAndValidate(c *fiber.Ctx, model interface{}) (interface{}, error) {
	contentType := c.Accepts(string(c.Request().Header.ContentType()))
	if !Contains(AcceptedContentType, contentType) {
		return nil, fmt.Errorf("invalid content type %s", contentType)
	}
	// Parse
	if deserializationError := c.BodyParser(model); deserializationError != nil {
		return nil, deserializationError
	}
	// Validate
	validationError := v.Struct(model)
	if validationError != nil {
		for _, err := range validationError.(validator.ValidationErrors) {
			return nil, fmt.Errorf("%s failed validation for %s. Enter a valid value", err.Field(), err.Tag())
		}
		return nil, validationError
	}
	return model, nil
}
