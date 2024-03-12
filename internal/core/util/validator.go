package util

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Validator struct {
}

func (v Validator) IsDate(fl validator.FieldLevel) bool {
	dateString := fl.Field().String()
	datePattern := `^\d{4}-\d{2}-\d{2}$`
	regex := regexp.MustCompile(datePattern)
	return regex.MatchString(dateString)
}

func NewValidator() *validator.Validate {
	v := Validator{}
	validate := validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("IsDate", v.IsDate)
	return validate
}
