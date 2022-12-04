package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/*
TODO: remake functionality got array of messages...
*/

type ValidationError struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidationErrorFromError(field string, err error) *ValidationError {
	return &ValidationError{
		FailedField: field,
		Value:       err.Error(),
	}
}

var validate = validator.New()

func ValidateStruct(structToValidate interface{}) []*ValidationError {
	var errors []*ValidationError
	err := validate.Struct(structToValidate)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ValidationError
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}

	return errors
}

func UnprocessableEntityResponse(ctx *fiber.Ctx, errors []*ValidationError) error {
	return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"messages": errors,
	})
}
