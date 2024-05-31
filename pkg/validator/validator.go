package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
}

type ValidationError struct {
	Field string `json:"field"`
	Msg   string `json:"msg"`
}

func ValidateStruct(s interface{}) []ValidationError {
	var errs []ValidationError
	err := Validate.Struct(s)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errs = make([]ValidationError, len(ve))
			for i, fe := range ve {
				errs[i] = ValidationError{strings.ToLower(fe.Field()), msgForTag(fe.Tag())}
			}
		}
	}
	return errs
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return ""
}
