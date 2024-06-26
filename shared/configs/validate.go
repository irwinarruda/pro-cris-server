package configs

import (
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/irwinarruda/pro-cris-server/shared/models"
)

var validate *validator.Validate

type Validate = *validator.Validate

func GetValidate(loginProviders []string) Validate {
	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("login_provider", ValidateLoginProvider(loginProviders))
		validate.RegisterValidation("weekday", ValidateWeekDay)
		validate.RegisterValidation("required_ifid", ValidateWeekDay)
		validate.RegisterValidation("weekday_ifid", ValidateWeekDay)
	}
	return validate
}

func ValidateLoginProvider(loginProviders []string) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		input := fl.Field().String()
		return slices.Contains(loginProviders, input)
	}
}

func ValidateWeekDay(fl validator.FieldLevel) bool {
	weekDays := models.GetWeekDays()
	input := fl.Field().String()
	return slices.Contains(weekDays, input)
}

func ValidateRequiredIfID(fl validator.FieldLevel) bool {
	id := fl.Field().FieldByName("ID")
	if id.IsNil() {
		return fl.Field().IsNil()
	}
	return true
}

func ValidateWeekdayIfID(fl validator.FieldLevel) bool {
	id := fl.Field().FieldByName("ID")
	if id.IsNil() {
		return ValidateWeekDay(fl)
	}
	return true
}
