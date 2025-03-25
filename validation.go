package validation

import (
	"errors"
	"strings"

	"github.com/go-playground/validator"
)

// validateRequest memvalidasi struct request dan mengembalikan error jika ada
func ValidateRequest(req interface{}) error {
	var validate = validator.New()
	if err := validate.Struct(req); err != nil {
		var errorMessages []string
		for _, err := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, formatValidationError(err))
		}
		return errors.New(strings.Join(errorMessages, ", "))
	}
	return nil
}

// formatValidationError membuat pesan error lebih spesifik
func formatValidationError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " tidak boleh kosong"
	case "email":
		return fe.Field() + " harus berupa email yang valid"
	case "min":
		return fe.Field() + " harus memiliki minimal " + fe.Param() + " karakter"
	case "max":
		return fe.Field() + " tidak boleh lebih dari " + fe.Param() + " karakter"
	default:
		return fe.Field() + " tidak valid"
	}
}
