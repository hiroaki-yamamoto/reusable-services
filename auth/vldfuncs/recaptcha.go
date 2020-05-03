package vldfuncs

import (
	"github.com/go-playground/validator/v10"
)

// Recaptcha is a validation function to check recaptcha text.
func Recaptcha(secret string) validator.Func {
	return func(fl validator.FieldLevel) bool {
	}
}
