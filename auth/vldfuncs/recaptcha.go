package vldfuncs

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/hiroaki-yamamoto/recaptcha"
)

// Recaptcha is a validation function to check recaptcha text.
func Recaptcha(secret string) validator.Func {
	return func(fl validator.FieldLevel) bool {
		addrFldValue, addrKind, _, success := fl.GetStructFieldOK2()
		if !success || addrKind != reflect.String {
			return false
		}
		respFldVal := fl.Field()
		if respFldVal.Kind() != reflect.String {
			return false
		}
		recap := recaptcha.New(secret)
		resp, err := recap.Check(addrFldValue.String(), respFldVal.String())
		if err != nil {
			return false
		}
		return resp.Success
	}
}
