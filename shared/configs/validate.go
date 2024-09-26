package configs

import (
	"slices"

	"github.com/go-playground/validator/v10"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

var validate *validator.Validate

type Validate = *validator.Validate

func GetValidate(loginProviders []string, paymentStyle []string, paymentType []string, settlementStyle []string) Validate {
	if validate == nil {
		validate = validator.New()
		validate.RegisterValidation("login_provider", ValidateEnum(loginProviders))
		validate.RegisterValidation("payment_style", ValidateEnum(paymentStyle))
		validate.RegisterValidation("payment_type", ValidateEnum(paymentType))
		validate.RegisterValidation("settlement_style", ValidateEnum(settlementStyle))
		validate.RegisterValidation("weekday", ValidateEnum(utils.Map(models.GetWeekDays(), func(m models.WeekDay, _ int) string {
			return m.String()
		})))
		validate.RegisterValidation("gender", ValidateEnum(models.GetGender()))
		validate.RegisterValidation("required_ifid", ValidateWeekDay)
		validate.RegisterValidation("weekday_ifid", ValidateWeekDay)
	}
	return validate
}

func ValidateEnum(enums []string) func(fl validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		input := fl.Field().String()
		return slices.Contains(enums, input)
	}
}

func ValidateWeekDay(fl validator.FieldLevel) bool {
	weekDays := models.GetWeekDays()
	input := fl.Field().String()
	return slices.Contains(weekDays, models.ToWeekDay(input))
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
