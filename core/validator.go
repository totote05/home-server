package core

import (
	"reflect"
	"time"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
		validate.RegisterCustomTypeFunc(ValidateTime, time.Time{})
	}

	return validate
}

func ValidateTime(field reflect.Value) interface{} {
	if value, ok := field.Interface().(time.Time); ok {
		if !value.IsZero() {
			return value
		}
	}
	return nil
}
