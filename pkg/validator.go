package pkg

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func (v *Validator) Validate(in any) error {
	if err := v.validator.Struct(in); err != nil {
		var validationErrors []string

		for _, err := range err.(validator.ValidationErrors) {
			field := err.Field()
			tag := err.Tag()

			switch tag {
			case "required":
				validationErrors = append(validationErrors, fmt.Sprintf("field '%s' is required", field))
			case "email":
				validationErrors = append(validationErrors, fmt.Sprintf("field '%s' must be a valid email", field))
			case "min":
				validationErrors = append(validationErrors, fmt.Sprintf("field '%s' must be at least %s", field, err.Param()))
			case "max":
				validationErrors = append(validationErrors, fmt.Sprintf("field '%s' must be at most %s", field, err.Param()))
			case "len":
				validationErrors = append(validationErrors, fmt.Sprintf("field '%s' must be exactly %s characters", field, err.Param()))
			case "uuid":
				validationErrors = append(validationErrors, fmt.Sprintf("field '%s' must be a valid UUID", field))
			default:
				validationErrors = append(validationErrors, fmt.Sprintf("field '%s' validation failed for '%s'", field, tag))
			}
		}

		return errors.New(strings.Join(validationErrors, "; "))
	}
	return nil
}
