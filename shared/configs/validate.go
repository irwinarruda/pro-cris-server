package configs

import (
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/irwinarruda/pro-cris-server/shared/models"
)

var validate *validator.Validate

type Validate = *validator.Validate

func GetValidate(validations map[string][]string) Validate {
	if validate == nil {
		validate = validator.New()
		for key, enums := range validations {
			validate.RegisterValidation(key, ValidateEnum(enums))
		}
		validate.RegisterValidation("weekday", ValidateEnum(models.GetWeekDaysString()))
		validate.RegisterValidation("gender", ValidateEnum(models.GetGenderString()))
		validate.RegisterValidation("required_ifid", ValidateWeekDay)
		validate.RegisterValidation("weekday_ifid", ValidateWeekDay)
	}
	return validate
}

func ValidateEnum(enums []string) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		input := fl.Field().String()
		if input == "" {
			return true
		}
		return slices.Contains(enums, input)
	}
}

func ValidateWeekDay(fl validator.FieldLevel) bool {
	weekDays := models.GetWeekDaysString()
	input := fl.Field().String()
	if input == "" {
		return true
	}
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
