package configs

import (
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/irwinarruda/pro-cris-server/shared/entities"
)

var validate *validator.Validate

type Validate = *validator.Validate

func GetValidate() Validate {
	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("weekday", ValidateWeekDay)
	}
	return validate
}

func ValidateWeekDay(fl validator.FieldLevel) bool {
	weekDays := entities.GetWeekDays()
	input := fl.Field().String()
	return slices.Contains(weekDays, input)
}
