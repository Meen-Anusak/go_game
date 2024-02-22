package helper

import (
	"github.com/go-playground/validator/v10"
)

var Validator = validator.New()

func Validate(data interface{}) []*ValidateError {
	var errors []*ValidateError
	err := Validator.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidateError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return errors
	}
	return nil
}

type ValidateError struct {
	Field string
	Tag   string
	Value string
}
